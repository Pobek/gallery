package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User - An object representing the a user from the database
type User struct {
	gorm.Model
	FirstName string `grom:"size:100;not null" json:"firstname"`
	LastName  string `grom:"size:100;not null" json:"lastname"`
	Username  string `grom:"size:100;not null" json:"username"`
	Password  string `grom:"size:100;not null" json:"password"`
}

// HashPassword - hashes password from user input
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // 14 is the cost for hasing
	return string(bytes), err
}

// CheckPasswordHash - checks if password equals to give hash
func CheckPasswordHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password incorrect")
	}
	return nil
}

// BeforeSave - hashes the user's password
func (user *User) BeforeSave() error {
	password := strings.TrimSpace(user.Password)
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// Prepare - strips user input of any whitespace
func (user *User) Prepare() {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Username = strings.TrimSpace(user.Username)
}

// Validate - validates user input
func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if user.Username == "" {
			return errors.New("Username is required to login")
		}
		if user.Password == "" {
			return errors.New("Password is required to login")
		}
		return nil
	default:
		if user.Username == "" {
			return errors.New("Username is required")
		}
		if user.Password == "" {
			return errors.New("Password is required")
		}
		if user.FirstName == "" {
			return errors.New("First name is required")
		}
		if user.LastName == "" {
			return errors.New("Last name is required")
		}
		return nil
	}
}

// SaveUser - adds user to the database
func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	// Debug a single operation, show detailed log for this operation
	err = db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// GetUser - returns a user based on Username
func (user *User) GetUser(db *gorm.DB) (*User, error) {
	account := &User{}
	if err := db.Table("users").Where("username = ?", user.Username).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

// GetAllUsers - returns all the users in the database
func GetAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	if err := db.Table("users").Find(&users).Error; err != nil {
		return &[]User{}, err
	}
	return &users, nil
}
