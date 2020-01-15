package main

import (
	"fmt"
	"querybuilder/db"
)

func main() {
	// Create a db object
	db1, err1 := db.New(db.MYSQL)
	if err1 != nil {
		fmt.Println("Error1 in db1: "+err1.Error())
		return
	}

	// Set a table name
	db1.Table("users")

	users1 := db1.Get() // Fetch all the users from users table
	fmt.Println(users1)

	user1 := db.Find(101) // Fetch a user with id 101
	fmt.Println(user1)

	remove_user101 := db1.Query("delete where id=101")
	fmt.Println(remove_user101)

	db2, err2 := db.New(db.MYSQL)
	if err2 != nil {
		fmt.Println("Error2 in db2: " + err2.Error())
		return
	}

	users2 := db2.Get() 
	fmt.Println(users2)
}