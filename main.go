package main

import (
	"database/sql"
	"fmt"
	"strings"
)

type DataBase struct {
	Source *sql.DB
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
var isDev bool = false
var LoggedUser User
var DBM DataBase

func main() {
	DBM.Source = InitDb()
	if !isDev {
		LoggedUser = User{Username: "", Password: ""}
		authorized := initialize(&LoggedUser)
		if !authorized {
			fmt.Println("SORRY YOU ARE NOT AUTHORIZED")
		}
	} else {
		LoggedUser = getAuthorizedUser("admin")
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
