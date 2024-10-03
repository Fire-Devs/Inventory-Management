package repository

import (
	"InventoryManagement/database"
	"InventoryManagement/models"
	"context"
	"github.com/jackc/pgx/v5"
)

func AddRoles(role *models.Role) error {
	conn, err := database.Connect()
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(),
		"INSERT INTO roles (name, permissions) VALUES ($1, $2)",
		role.Name, role.Permissions)

	if err != nil {
		return err
	}

	return nil
}

func FetchRoles(name string) ([]models.Role, error) {
	conn, err := database.Connect()
	if err != nil {
		return nil, err
	}

	query := "SELECT id, name, permissions FROM roles"

	var rows pgx.Rows

	if name != "" {
		query += " WHERE name = $1"
		rows, err = conn.Query(context.Background(), query, name)
		if err != nil {
			return nil, err
		}

	} else {
		rows, err = conn.Query(context.Background(), query)
		if err != nil {
			return nil, err
		}

	}

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name, &role.Permissions)

		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func UpdateRoles(role *models.Role) error {
	conn, err := database.Connect()
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), "UPDATE roles SET name = $1, permissions = $2 WHERE id = $3", role.Name, role.Permissions, role.ID)
	if err != nil {
		return err
	}

	return nil
}

func FetchPermissionfromUser(user string) ([]string, error) {
	conn, err := database.Connect()
	if err != nil {
		return nil, err
	}

	var permissions []string
	err = conn.QueryRow(context.Background(), "SELECT r.permissions FROM users u JOIN roles r ON u.role = r.id WHERE u.email = $1", user).Scan(&permissions)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
