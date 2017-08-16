package main

import (
	"fmt"
)

var currentId int
var users Users

func init() {
	RepoCreateUser(User{Name: "John Do"})
	RepoCreateUser(User{Name: "Willy Wonka"})
	RepoCreateUser(User{Name: "Dr Yo"})
}

func RepoGetUser(id int) User {
	for _, user := range users {
		if user.Id == id {
			return user
		}
	}

	// return empty User if not found
	return User{}
}

func RepoCreateUser(user User) User {
	currentId += 1

	user.Id = currentId
	users = append(users, user)

	return user
}

func RepoUpdateUser(user User) User {
	for i := 0; i < len(users); i++ {
		if user.Id == users[i].Id {
			users[i] = user
		}
	}

	return user
}

func RepoDeleteUser(id int) error {
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d to delete", id)
}
