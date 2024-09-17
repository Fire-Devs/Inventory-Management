package handler

import (
	"InventoryManagement/config"
	"InventoryManagement/models"
	"InventoryManagement/repository"
	"InventoryManagement/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var conf config.Config

func init() {
	conf = config.LoadConfig()
}

func Login(c fiber.Ctx) error {

	user := new(models.User)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	name, email, password, err := repository.GetUserByEmailOrName(user.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.CheckPassword(user.Password, password); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	t, err := token.SignedString([]byte(conf.Jwt.Secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		MaxAge:   60 * 60 * 48,
		Secure:   false,
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"success": "Logged IN",
	})

}

func ParseJWT(cookie string) error {

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Jwt.Secret), nil
	})

	if err != nil {
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	data := claims["email"].(string)

	_, _, _, err2 := repository.GetUserByEmailOrName(data)
	if err2 != nil {
		return err
	}

	return nil
}

func Register(c fiber.Ctx) error {
	user := new(models.User)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	name, email, _, _ := repository.GetUserByEmailOrName(user.Email)
	if email != "" && name != "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err2 := repository.CreateUser(user.Name, user.Email, hashedPassword)
	if err2 != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err2.Error(),
		})
	}

	randomtoken := utils.GenerateARandomString(20)
	err = utils.SendEmail("http://localhost:8080/auth/verify?token="+randomtoken, user.Email)
	err = repository.SetUserToken(randomtoken, user.Email)

	return c.JSON(fiber.Map{
		"success": "Check Mail",
	})

}

func VerifyToken(c fiber.Ctx) error {
	token := c.Query("token")
	if token == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Token not found",
		})
	}

	_, err := repository.GetUserToken(token)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Token Invalid or Expired",
		})
	}

	err2 := repository.Verifyuser(token)
	if err2 != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err2.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "User verified",
	})

}
