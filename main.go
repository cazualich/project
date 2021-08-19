package main

import (
	"fmt"
)


type User struct {
	ID int
	Name string
	Email string
}

func (u *User) changeName(name string) {
	u.Name = name
}

func (u User) getName() string {
	return u.Name
}

func (u User) String() string {
	return fmt.Sprint("User name: ", u.Name)
}

func main() {
	user1 := User {
		ID: 1,
		Name: "Misha",
		Email: "abcdef",
	}
	user2 := User {
		Name: "Ivan",
		Email: "qwerty",
	}
	users := []User{
		user1,
		user2,
	}
	for _, user := range users {
		fmt.Println(user.String())
	}
}
