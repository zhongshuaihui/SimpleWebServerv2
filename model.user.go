package main

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
