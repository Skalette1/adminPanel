package models

import (
	"database/sql"
	"errors"
	//_ "github.com/Skalette1/adminPanel"
)

type Role struct {
	ID         int    `jsonapi:"primary,roleApi"`
	Username   string `jsonapi:"attr,username"`
	Permission string `jsonapi:"attr,permission"`
	CreatedAt  string `jsonapi:"attr, createdAt"`
	UpdatedAt  string `jsonapi:"attr, updatedAt"`
}

func CreateRolesTable(db *sql.DB) error {
	schema := `
		CREATE TABLE IF NOT EXISTS roles(
		    ID INTEGER SERIAL PRIMARY KEY,
		    username VARCHAR(255) NOT NULL DEFAULT "",
		    permissions VARCHAR(255) NOT NULL DEFAULT "",
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
`
	if _, err := db.Exec(schema); err != nil {
		return errors.New("create roles table: " + err.Error())
	}
	return nil
}
