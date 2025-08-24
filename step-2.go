package main

import "fmt"

func main() {
	fmt.Println("This is one", 1)
	fmt.Println("This is two", 2)
	fmt.Println("This is decimal", 3.5)

	fmt.Println(len("hokeyyy"))
	fmt.Println("hookeyyy"[0])
	fmt.Println("hookeyyy"[1])

	// variable access
	var name string
	name = "first name"
	fmt.Println("name=", name)

	name = "last name"
	fmt.Println("name=", name)

	// simplified access variable
	anotherName := "cool"
	fmt.Println("anothername=", anotherName)

	anotherName = "another cool"
	fmt.Println("anothername=", anotherName)

	// multiple variable declaration
	var (
		cardName   = "name"
		cardNumber = 123
	)
	fmt.Println("card", cardName, cardNumber)

	// constants
	const (
		constName = "name"
		constId   = 123
	)
	fmt.Println("constants", constName, constId)

	// "type" data
	type NoKTP string

	var ktpPersonOne string = "12321231"
	var ktpPersonTwo NoKTP = NoKTP("5454645")

	fmt.Println("type", ktpPersonOne, ktpPersonTwo)

	// array
	var namesWhosJoin = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("array", namesWhosJoin)
}
