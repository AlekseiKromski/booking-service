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
