package main

import "fmt"

func checkUser(username string, password string) (bool, User) {
	rows, err := DBM.Source.Query("SELECT User_.id, User_.username, User_.password, Role_.name FROM User_ INNER JOIN Role_ ON User_.role_id = Role_.id WHERE Username = $1 AND Password = $2 LIMIT 1 ", username, password)
	if err != nil {
		fmt.Println(err)
	}

	user := User{Username: "", Password: "", Role: Role{Name: ""}}
	if rows != nil {
		for rows.Next() {
			err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role.Name)
			if err != nil {
				panic(err)
			}
			fmt.Println("\n", user)
		}
		user.Role.Actions = getActions(user.Role.Name)
		return true, user

	} else {
		return false, User{}
	}
	return false, User{}

}
func adminAddNewCustomer() {
	customer := User{Username: "", Password: "", Role: Role{Name: "customer", Actions: nil}}

infinity:
	for {
		fmt.Print("Username: ")
		fmt.Scanln(&customer.Username)

		rows, err := DBM.Source.Query("SELECT count(*) FROM User_ WHERE Username = $1 LIMIT 1", customer.Username)
		if err != nil {
			fmt.Println(err)
		}
		var count int
		for rows.Next() {
			rows.Scan(&count)
		}
		if count != 0 {
			fmt.Println("This customer already exists, try again")
			continue infinity
		}

		_, err = DBM.Source.Exec("INSERT INTO User_ (username, password, role_id) VALUES ($1,$2,$3)", customer.Username, customer.Username, 2)
		customer.Password = customer.Username

		if err != nil {
			fmt.Println(err)
		}
		break infinity

	}

	fmt.Printf("New customer was created: \n - [Username: %s]\n - [Password: %s]\n", customer.Username, customer.Password)
}
func adminListOfCustomers() {
	rows, err := DBM.Source.Query("SELECT Username, Password FROM User_ WHERE role_id = 2")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var username string
		var password string
		rows.Scan(&username, &password)
		fmt.Printf("\n - [Username: %s]\n - [Password: %s]\n\n", username, password)

	}
}
func customerAddBook() {
	newBook := Book{BookNumber: "", DateStart: "", DateEnd: "", Notice: "", Status: "", User: User{}}

	fmt.Print("Book number: ")
	fmt.Scanln(&newBook.BookNumber)

	fmt.Print("Status: ")
	fmt.Scanln(&newBook.Status)

	fmt.Print("Notice: ")
	fmt.Scanln(&newBook.Notice)

	fmt.Print("Book start date: ")
	fmt.Scanln(&newBook.DateStart)

	fmt.Print("Book end date: ")
	fmt.Scanln(&newBook.DateEnd)

	newBook.User = *getUser()

	_, err := DBM.Source.Exec("INSERT INTO book (booknumber, datestart,dateend,status,notice,user_id) VALUES ($1,$2,$3,$4,$5,$6)", newBook.BookNumber, newBook.DateStart, newBook.DateEnd, newBook.Status, newBook.Notice, newBook.User.Id)
	if err != nil {
		fmt.Println(err)
	}
	printBook(newBook)

}
func userCheckBookingStatus() {
	book := Book{BookNumber: "", Status: "", Notice: "", DateStart: "", DateEnd: ""}
	fmt.Print("Booking number: ")
	fmt.Scanln(&book.BookNumber)
	rows, err := DBM.Source.Query("SELECT BookNumber, Status, Notice, datestart, dateEnd FROM Book WHERE booknumber = $1", book.BookNumber)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&book.BookNumber, &book.Status, &book.Notice, &book.DateStart, &book.DateEnd)
		printBook(book)
	}
}
func customerListOfBookings() {
	book := Book{BookNumber: "", Status: "", Notice: "", DateStart: "", DateEnd: ""}
	rows, err := DBM.Source.Query("SELECT BookNumber, Status, Notice, datestart, dateEnd FROM Book")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&book.BookNumber, &book.Status, &book.Notice, &book.DateStart, &book.DateEnd)
		printBook(book)
	}
}
func printBook(book Book) {
	fmt.Printf("\n - [BOOK NUMBER: %s]\n - [Book status: %s]\n - [Book notice: %s]\n - [Book start date: %s]\n - [Book end date: %s]\n",
		book.BookNumber, book.Status, book.Notice, book.DateStart, book.DateEnd)
}
