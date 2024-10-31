package models

type User struct {
	ID       int
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	ID       int
	Email    string `json:"email" validate:"required_without=Name"`
	Name     string `json:"name" validate:"required_without=Email"`
	Password string `json:"password" validate:"required"`
}
