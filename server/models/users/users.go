package users

import (
	"fmt"

	"github.com/kunalpowar/gopggqlauth/server/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id" pg:"id"`
	Username string `json:"name" pg:"username"`
	Password string `json:"password" pg:"password"`
}

func (user *User) Create() error {
	model := User{Username: user.Username}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("users: could not hash password: %v", err)
	}

	model.Password = hashedPassword

	if _, err := db.DB.Model(&model).Insert(); err != nil {
		return fmt.Errorf("users: could not insert user %v: %v", model, err)
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
	if err := db.DB.Model(&u).Where("username = ?", username).Select(); err != nil {
		return 0, fmt.Errorf("users: could not get user for username %q: %v", username, err)
	}

	return u.ID, nil
}

func (user *User) Authenticate() (bool, error) {
	var u User
	if err := db.DB.Model(&u).Where("username = ?", user.Username).Select(); err != nil {
		return false, fmt.Errorf("users: could not get user for username %q: %v", user.Username, err)
	}

	return CheckPasswordHash(user.Password, u.Password), nil
}
