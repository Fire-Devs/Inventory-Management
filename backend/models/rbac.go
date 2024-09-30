package models

import "github.com/go-playground/validator/v10"

var Permissions = []string{
	"create:inventory",
	"read:inventory",
	"update:inventory",
	"delete:inventory",
	"create:category",
	"read:category",
	"update:category",
	"delete:category",
	"create:supplier",
	"read:supplier",
	"update:supplier",
	"delete:supplier",
	"create:user",
	"read:user",
	"update:user",
	"delete:user",
	"create:role",
	"read:role",
	"update:role",
	"delete:role",
}

type Role struct {
	ID          int
	Name        string   `json:"name" validate:"required"`
	Permissions []string `json:"permissions" validate:"required,validPermissions"`
}

func ValidPermissions(fl validator.FieldLevel) bool {
	permissions := fl.Field().Interface().([]string)
	for _, permission := range permissions {
		if !contains(Permissions, permission) {
			return false
		}
	}
	return true
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
