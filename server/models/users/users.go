package users

import (
	"fmt"

	"github.com/kunalpowar/gopggqlauth/server/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"name" db:"username"`
	Password string `json:"password" db:"password"`
}

func (user *User) Create() error {
	u := User{Username: user.Username}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("users: could not hash password: %v", err)
	}
	u.Password = hashedPassword

	if _, err := db.Instance.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", u.Username, u.Password); err != nil {
		return fmt.Errorf("users: could not insert user %v: %v", u, err)
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIdByUsername(username string) (int, error) {
	var u User
	if err := db.Instance.Get(&u, "SELECT * FROM users WHERE username=$1", username); err != nil {
		return 0, fmt.Errorf("users: could not get user for username %q: %v", username, err)
	}

	return u.ID, nil
}

func (user *User) Authenticate() (bool, error) {
	var u User
	if err := db.Instance.Get(&u, "SELECT * FROM users WHERE username=$1", user.Username); err != nil {
		return false, fmt.Errorf("users: could not get user for username %q: %v", user.Username, err)
	}

	return CheckPasswordHash(user.Password, u.Password), nil
}
