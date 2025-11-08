package repository

import (
	"database/sql"

	"github.com/Skalette1/adminPanel/internal/models"
)

type RoleRepository struct {
	DB *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) Create(role models.Role) (int, error) {
	var id int
	err := r.DB.QueryRow(`
		INSERT INTO roles (username, permission)
		VALUES ($1, $2)
		RETURNING id;
	`, role.Username, role.Permission).Scan(&id)

	return id, err
}

func (r *RoleRepository) GetByID(id int) (*models.Role, error) {
	var role models.Role
	err := r.DB.QueryRow(`
		SELECT id, username, permission, created_at, updated_at
		FROM roles WHERE id=$1
	`, id).Scan(
		&role.ID,
		&role.Username,
		&role.Permission,
		&role.CreatedAt,
		&role.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) GetAll() ([]models.Role, error) {
	rows, err := r.DB.Query(`SELECT id, username, permission, created_at, updated_at FROM roles`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role

	for rows.Next() {
		var role models.Role
		if err := rows.Scan(
			&role.ID, &role.Username, &role.Permission, &role.CreatedAt, &role.UpdatedAt,
		); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (r *RoleRepository) Update(id int, role models.Role) error {
	res, err := r.DB.Exec(`
		UPDATE roles
		SET username = $1, permission = $2, updated_at = now()
		WHERE id = $3
	`, role.Username, role.Permission, id)
	if err != nil {
		return err
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *RoleRepository) Delete(id int) error {
	res, err := r.DB.Exec(`DELETE FROM roles WHERE id=$1`, id)
	if err != nil {
		return err
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}
