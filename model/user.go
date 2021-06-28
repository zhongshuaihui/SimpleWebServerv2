package model

import (
	"errors"
	"simplewebserverv2/middleware"
	"strings"
)

type user struct {
	User_name string `json:"user_name"`
	Password  string `json:"password"`
}

// v1: the user store in disk
// var userList = []user{
// 	{User_name: "user1", Password: "123"},
// 	{User_name: "user2", Password: "123"},
// 	{User_name: "user3", Password: "123"},
// }

// v2: the user store in database
var userList []user

func IsUserValid(username string, password string) bool {
	if len(userList) == 0 {
		ReadUserList()
	}
	for _, user := range userList {
		if user.User_name == username && user.Password == password {
			return true
		}
	}
	return false
}

func RegisterNewUser(username string, password string) (err error) {
	if strings.TrimSpace(password) == "" {
		return errors.New("the password can not be empty")
	}
	if !IsUserNameAvaliable(username) {
		return errors.New("this name is already registered")
	}

	// u := user{User_name: username, Password: password}
	// userList = append(userList, u)

	stmt, err := middleware.Db.Prepare("insert into user values (?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, password)
	if err != nil {
		return
	}
	return nil
}

func IsUserNameAvaliable(username string) bool {
	if len(userList) == 0 {
		ReadUserList()
	}
	for _, user := range userList {
		if username == user.User_name {
			return false
		}
	}
	return true
}

func ReadUserList() (err error) {
	rows, err := middleware.Db.Query("select * from user")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		rows.Scan(&u.User_name, &u.Password)
		userList = append(userList, u)
	}

	return
}
