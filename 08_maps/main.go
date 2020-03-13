package main

import "fmt"

func main() {
	//Define a map
	//emails := make(map[string]string)

	// Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Rick"] = "rick@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add values
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	//Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
