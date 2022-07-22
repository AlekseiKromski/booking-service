package main

import (
	"fmt"
	"os"
)

func getActions(role string) map[string]func() {
	actions := make(map[string]func())
	actions["0. Exit"] = func() {
		fmt.Println("PROCESS OF EXIT")
		os.Exit(3)
	}
	if role == "admin" {
		actions["1. Add new customer"] = func() {
			customer := User{Username: "", Password: "", Role: Role{Name: "customer", Actions: nil}}
			adminAddNewCustomer(&customer, &users)
			users = append(users, customer)
		}
		actions["2. Show list of customers"] = func() {
			adminListOfCustomers(users)
		}
	} else if role == "customer" {
		actions["1. Add book"] = func() {
			newBook := Book{BookNumber: "", DateStart: "", DateEnd: "", Notice: "", Status: "", User: User{}}
			customerAddBook(bookings, &newBook)
			bookings = append(bookings, newBook)
		}
		actions["2. Show list of book"] = func() {
			customerListOfBookings(bookings)
		}
	} else if role == "user" {
		actions["1. check booking status"] = func() {
			fmt.Println("PROCESS OF ADDING")
		}
	}
	return actions

}

var users []User = []User{User{
	Username: "admin",
	Password: "admin",
	Role:     Role{Name: "admin", Actions: nil},
}}

var bookings []Book = []Book{}

func checkUser(username string, password string) (bool, User) {
	for _, user := range users {
		if user.Username == username && password == user.Password {
			user.Role.Actions = getActions(user.Role.Name)
			return true, user
		}
	}
	return false, User{}
}

//For testing
func getAuthorizedUser(role string) User {
	user := User{
		Username: "admin",
		Password: "admin",
		Role:     Role{Name: role, Actions: nil},
	}

	user.Role.Actions = getActions(user.Role.Name)
	return user
}
