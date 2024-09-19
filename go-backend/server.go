package main

import "fmt"

func main() {
	DatabaseConnection()

	users, err := GetUsers()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(users)
	}

	user, err := GetUserById(3)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}

	err = AddUser(UserData{
		name: "Mathias Kornidesz",
	})

	if err != nil {
		fmt.Println(err)
	}

	users, err = GetUsers()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(users)
	}
}
