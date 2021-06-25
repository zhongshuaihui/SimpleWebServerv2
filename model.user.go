package main

import (
	"errors"
	"strings"
)

type user struct {
	User_name string `json:"user_name"`
	Password  string `json:"password"`
}

var userList = []user{
	{User_name: "user1", Password: "123"},
	{User_name: "user2", Password: "123"},
	{User_name: "user3", Password: "123"},
}

func isUserValid(username string, password string) bool {
	for _, user := range userList {
		if user.User_name == username && user.Password == password {
			return true
		}
	}
	return false
}

func registerNewUser(username string, password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("the password can not be empty")
	}
	if !isUserNameAvaliable(username) {
		return errors.New("this name is already registered")
	}

	u := user{User_name: username, Password: password}
	userList = append(userList, u)
	return nil
}

func isUserNameAvaliable(username string) bool {
	for _, user := range userList {
		if username == user.User_name {
			return false
		}
	}
	return true
}
