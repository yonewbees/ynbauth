package models

import (
	"database/sql"
	"ynbauth/database"
	"errors"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func CreateUser( username, fullName, email, password string) error {
	
	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Check if the user already exists
	existingUser, _ := FindUser(email)
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	// Insert the user into the database
	query := `INSERT INTO users (username, full_name, email, password) VALUES ($1, $2, $3, $4)`
	_, err = database.DB.Exec(query, username, fullName, email, string(hashedPassword))
	if err != nil {
		return errors.New("failed to insert user into the database")
	}

	return nil
}


// FindUser checks if a user exists by email or username and returns the user object if found.
func FindUser(identifier string) (*User, error) {
	var user User

	query := `
		SELECT id, username, full_name, email, password
		FROM users 
		WHERE username = $1 OR email = $1
	`

	err := database.DB.QueryRow(query, identifier).Scan(
		&user.ID,
		&user.Username,
		&user.FullName,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}