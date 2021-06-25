package service

import (
	gorm "Auth/internal/database/gorm/user"
	"Auth/internal/user/model"
	tok "Auth/internal/user/service/token"
	"fmt"
	"log"
	//"Auth/internal/user"
)

func Signup(user model.User) (model.User, error) {
	//Hash Password
	password := user.Password
	hash, err := tok.HashPassword(password)
	// Hash password to database and gives token
	user.Password = hash
	fmt.Print(hash)
	// use create user func from model
	user.Role = "member"
	usr, err := gorm.Createuser(&user)
	if err != nil {
		return model.User{}, err
	}
	return usr, err
}

func Signin(user *model.User) (*tok.TokenDetails, error) {
	usr, _ := gorm.Getusersbyname(user.Username)
	log.Println("User name is :", user.Username)
	password := user.Password
	//hash, err := HashPassword(password)

	if !tok.CheckPasswordHash(password, usr.Password) {
		return nil, nil
	}

	//token, _ := CreateToken(&user)
	fmt.Println("testing..", &user)
	token, err := tok.CreateToken(&usr)
	// Get users from model
	// Check password == with existing password
	// execute token
	return token, err
}
