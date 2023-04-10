package models

import (
	"errors"
	"mini-blog-go/utils/token"

	"golang.org/x/crypto/bcrypt"
)

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username" gorm:"not null; unique"`
	Password string `json:"password" gorm:"not null; unique"`
	Posts    []Post `gorm:"foreignKey:UserID"`
}

func (u *User) SaveUser() (*User, error) {
	var err error
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(passwordHashed)
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

//func (u *User) BeforeSave() error {
//	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
//	if err != nil {
//		return err
//	}
//	u.Password = string(passwordHashed)
//
//	return nil
//}

func PasswordVerify(password string, passwordHashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}

func LoginVerify(username string, password string) (string, error) {
	var err error

	user := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = PasswordVerify(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}
