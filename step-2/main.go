package main

import (
	"fmt"
)

func greeting(name string) string {
	return "Hello " + name
}

func main() {
	var age int = 20
	country := "indonesia"

	var nonValueVar string

	fmt.Println("default value of nonValuedVar is", nonValueVar)

	if age > 18 {
		fmt.Println("you are an adult")
	} else {
		fmt.Println("still a kid")
	}

	for i := range 3 {
		fmt.Println("Loop in", i)
	}

	country = "Japan" // re assign the value
	message := greeting("adib")
	fmt.Println("Greeting func:", message)

	fmt.Printf("Country: %s | Age: %d\n", country, age)
}
