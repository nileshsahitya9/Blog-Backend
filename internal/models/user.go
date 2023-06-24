package models

import (
	"database/sql"
	"errors"

	"github.com/nileshsahitya9/Blogs-Backend/db"
	"github.com/nileshsahitya9/Blogs-Backend/internal/common"
)

func CreateUser(user *common.User) error {
	query := "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id"
	err := db.DB.QueryRow(query, user.Username, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(userID int) (*common.User, error) {
	query := "SELECT id, username, email, bio, name FROM users WHERE id = $1"
	row := db.DB.QueryRow(query, userID)

	var user common.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Bio, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user *common.User) error {
	query := "UPDATE users SET name=$1, bio=$2 WHERE id = $3"
	_, err := db.DB.Exec(query, user.Name, user.Bio, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func CheckEmailExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	err := db.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
