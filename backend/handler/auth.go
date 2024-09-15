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

	name, email, password, err := repository.GetUserByEmailOrName(user.Name)
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

func Register(c fiber.Ctx) error {
	user := new(models.User)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	t, err := token.SignedString([]byte(conf.Jwt.Secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
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

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		MaxAge:   60 * 60 * 48,
		Secure:   false,
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"success": "Cookie set",
	})

}
