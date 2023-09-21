package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// Variabel dengan tipe data
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("Hello %s %s!\n\n", firstName, lastName);

	// Variabel tanpa tipe data
	p2FirstName := "John"
	p2LastName := "Travolta"

	fmt.Printf("Hello %s %s!\n\n", p2FirstName, p2LastName)
}
