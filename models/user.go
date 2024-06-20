package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
}

func (user *User) SetPassword(password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return nil
}

func (user *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    return err == nil
}
