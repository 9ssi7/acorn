package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@doe.com",
	}

	fmt.Printf("%c\n", 'A')             // A             | Unicode code point
	fmt.Printf("%d\n", 10)              // 10            | Decimal
	fmt.Printf("%f\n", 3.14)            // 3.140000      | Floating point
	fmt.Printf("%9.2f\n", 3.14)         //      3.14     | Floating point with width and precision
	fmt.Printf("%e\n", 3.14)            // 3.140000e+00  | Scientific notation
	fmt.Printf("%t\n", true)            // true          | Boolean
	fmt.Printf("%s\n", "Hello, World!") // Hello, World! | String
	fmt.Printf("%v\n", 10)              // 10            | Default format
	fmt.Printf("%b\n", 10)              // 1010          | Binary
	fmt.Printf("%x\n", 10)              // a             | Hexadecimal

	// Bonus
	fmt.Printf("%v\n", person)  // {John Doe 30 john@doe.com} 				 | Default
	fmt.Printf("%+v\n", person) // {Name:John Doe Age:30 Email:john@doe.com} | Field names
}
