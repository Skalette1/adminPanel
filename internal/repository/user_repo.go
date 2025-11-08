package repository

import (
	"database/sql"
	"errors"
	"github.com/Skalette1/adminPanel/internal/models"
	"log"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) (int, error) {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return 0, errors.New("invalid input")
	}

	var id int
	err := r.DB.QueryRow(`
		INSERT INTO users (username, email, password, role_id, is_active, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		RETURNING id`,
		user.Username, user.Email, user.Password, user.RoleID, user.IsActive, time.Now(), time.Now(),
	).Scan(&id)

	if err != nil {
		log.Println("CreateUser error:", err)
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	var user models.User
	err := r.DB.QueryRow(`
		SELECT id, username, email, password, role_id, is_active, created_at, updated_at
		FROM users WHERE id=$1`, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.RoleID,
		&user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	if user.ID <= 0 {
		return errors.New("invalid id")
	}
	_, err := r.DB.Exec(`
		UPDATE users 
		SET username=$1, email=$2, password=$3, role_id=$4, is_active=$5, updated_at=$6
		WHERE id=$7`,
		user.Username, user.Email, user.Password, user.RoleID, user.IsActive, time.Now(), user.ID,
	)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	res, err := r.DB.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query(`
		SELECT id, username, email, password, role_id, is_active, created_at, updated_at
		FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.RoleID,
			&user.IsActive, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
