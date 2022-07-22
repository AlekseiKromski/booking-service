package main

import (
	"fmt"
)

func adminAddNewCustomer(customer *User, users *[]User) {

infinity:
	for {
		fmt.Print("Username: ")
		fmt.Scan(&customer.Username)

		for _, username := range *users {
			if username.Username == customer.Username {
				fmt.Println("This customer already exists, try again")
				continue infinity
			}
		}
		break infinity

	}

	customer.Password += customer.Username + "_BookingService"

	fmt.Printf("New customer was created: \n - [Username: %s]\n - [Password: %s]\n", customer.Username, customer.Password)
}

func adminListOfCustomers(users []User) {
	for _, user := range users {
		if user.Role.Name == "customer" {
			fmt.Printf("\n - [Username: %s]\n - [Password: %s]\n\n", user.Username, user.Password)
		}
	}
}

func customerAddBook(bookings []Book, newBook *Book) {
	fmt.Print("Book number: ")
	fmt.Scan(&newBook.BookNumber)

	fmt.Print("Status: ")
	fmt.Scan(&newBook.Status)

	fmt.Print("Notice: ")
	fmt.Scan(&newBook.Notice)

	fmt.Print("Book start date: ")
	fmt.Scan(&newBook.DateStart)

	fmt.Print("Book end date: ")
	fmt.Scan(&newBook.DateEnd)

	newBook.User = *getUser()

	printBook(*newBook)

}

func customerListOfBookings(bookings []Book) {
	var output string

	for _, book := range bookings {
		if book.User.Username == getUser().Username {
			printBook(book)
		}
	}

	fmt.Println(output)
}

func printBook(book Book) {
	fmt.Printf("\n - [BOOK NUMBER: %s]\n - [Book status: %s]\n - [Book notice: %s]\n - [Book start date: %s]\n - [Book end date: %s]\n",
		book.BookNumber, book.Status, book.Notice, book.DateStart, book.DateEnd)
}
