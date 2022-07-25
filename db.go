package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type connection struct {
	User     string
	Password string
	Dbname   string
	Host     string
}

func getActions(role string) map[string]func() {
	actions := make(map[string]func())
	actions["0. Exit"] = func() {
		fmt.Println("PROCESS OF EXIT")
		os.Exit(3)
	}
	if role == "admin" {
		actions["1. Add new customer"] = func() {

			adminAddNewCustomer()
		}
		actions["2. Show list of customers"] = func() {
			adminListOfCustomers()
		}
	} else if role == "customer" {
		actions["1. Add book"] = func() {
			customerAddBook()
		}
		actions["2. Show list of book"] = func() {
			customerListOfBookings()
		}
	} else if role == "user" {
		actions["1. check booking status"] = func() {
			userCheckBookingStatus()
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

func InitDb() *sql.DB {

	//init db connection and work with models
	connection := connection{User: "postgres", Password: "", Host: "localhost", Dbname: "booking-service"}
	connStr := "postgres://" + connection.User + ":" + connection.Password + "@" + connection.Host + "/" + connection.Dbname + "?sslmode=disable"
	println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		println(err)
	}

	fmt.Println("DATABASE CONNECT SUCCESSFULY")

	//START MIGRATION IF NOT EXISTS
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Role_(Id SERIAL PRIMARY KEY, Name VARCHAR)`)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS User_ (Id SERIAL PRIMARY KEY,Username VARCHAR, Password VARCHAR, Role_id int,
	CONSTRAINT FK_ROLE
	 FOREIGN KEY(Role_id)
	 REFERENCES Role_(Id))`)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Book (
		Id SERIAL PRIMARY KEY,
		BookNumber VARCHAR,
		DateStart varchar,
		DateEnd varchar,
		Status varchar,
		Notice varchar,
		User_id int,
		CONSTRAINT FK_USER
		  FOREIGN KEY(User_id)
		  REFERENCES User_(Id)
	)`)

	if err != nil {
		fmt.Println(err)
	}
	return db
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
