package main

import (
	"fmt"
	"strings"
)

//TASKS
/*
	1) User can check status of your booking
	2) Administrator can add new booking code
	3) Program collect information about user actions

	System should support:
	1) login / logout / registration
	2) Roles - (For administrative role user should enter code, that will be given by admin) (User / Customer / Admin)
*/
type Role struct {
	Name    string
	Actions map[string]func()
}

type User struct {
	Username string
	Password string
	Role     Role
}

type Book struct {
	BookNumber string
	DateStart  string
	DateEnd    string
	Status     string
	Notice     string
	User       User
}

func (u User) showMenu() {
	printTitle("MENU")
	for menuTitle, _ := range u.Role.Actions {
		fmt.Println(menuTitle)
	}
}

func (u User) waitAction() string {
	var action string
	fmt.Print("Action number: ")
	fmt.Scan(&action)
	return strings.Trim(action, " ")
}

var authAttempts int = 4
var isDev bool = true
var LoggedUser User

func main() {
	if !isDev {
		LoggedUser = User{Username: "", Password: ""}
		authorized := initialize(&LoggedUser)
		if !authorized {
			fmt.Println("SORRY YOU ARE NOT AUTHORIZED")
		}
	} else {
		LoggedUser = getAuthorizedUser("customer")
	}

infinity:
	for {
		LoggedUser.showMenu()
		action := LoggedUser.waitAction()
		actionRunned := false
		for menuTitle, roleAction := range LoggedUser.Role.Actions {
			if menuTitle[0:1] == action {
				roleAction()
				continue infinity
			}
		}

		if !actionRunned {
			fmt.Println("NOT VALID ACTION")
		}

	}
}
func printTitle(title string) {
	var output string
	for i := 0; i < 9; i++ {
		if i < 3 || i > 5 {
			output += "="
		} else if i == 3 || i == 5 {
			output += " "
		} else {
			output += title
		}
	}
	fmt.Println(output)
}
func initialize(unAuthorizedUser *User) bool {
	printTitle("BOOKING SERVICE")

	for i := 0; i < authAttempts; i++ {
		if i > 0 {
			fmt.Println("----------")
		}

		fmt.Print("Username: ")
		fmt.Scan(&unAuthorizedUser.Username)

		fmt.Print("Password: ")
		fmt.Scan(&unAuthorizedUser.Password)

		fmt.Print("Try to auth - ")
		if result, user := checkUser(unAuthorizedUser.Username, unAuthorizedUser.Password); result {
			fmt.Print("OK\n")
			*unAuthorizedUser = user
			return true
		} else {
			fmt.Print("SOMETHING IS WRONG\n")
		}
	}

	return false
}

func getUser() *User {
	return &LoggedUser
}
