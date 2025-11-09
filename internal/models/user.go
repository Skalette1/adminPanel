package models

import (
	"database/sql"
	"errors"
	//	_ "github.com/Skalette1/adminPanel"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleId    int    `json:"role_id"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func CreateUsersTable(db *sql.DB) error {
	schema := `
		CREATE TABLE IF NOT EXISTS roles(
		    ID INTEGER SERIAL PRIMARY KEY,
		    username VARCHAR(255) NOT NULL DEFAULT "",
		    email VARCHAR(255) NOT NULL DEFAULT "",
		    password VARCHAR(255) NOT NULL DEFAULT "",
		    role_id INTEGER NOT NULL DEFAULT 0,
		    is_active BOOLEAN NOT NULL DEFAULT FALSE,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
`
	if _, err := db.Exec(schema); err != nil {
		return errors.New("create roles table: " + err.Error())
	}
	return nil
}
