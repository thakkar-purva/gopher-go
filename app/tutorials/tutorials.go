package tutorials

import (
	"fmt"
)

func PrintsTutorial() {
	fmt.Println("Welcome to the Go tutorials!")
}

// Go uses capitalization as a simple and clear way to distinguish between exported and unexported identifiers.
//  This avoids the need for additional keywords like public or private and keeps the language syntax minimal and clean.

func PrintsVariable() {

	// variable name can have letters, digits, and underscores only first_name, first31 etc
	//invalid name can not start with a digit 31name
	// variable name can not be a keyword like func, package, var, etc

	fmt.Println("All about golang variables !! ")

	var name string = "Purva Thakkar"
	fmt.Println("Hi", name)

	var age uint = 32
	fmt.Println("Age is ", age)
}

func PrintImplicitvsExplicit() {

	// implicit type declaration

	var number = 10            // implicit type
	var number2 int = 20       // explicit type
	fmt.Printf("%T\n", number) // print type of variable

	fmt.Printf("%T\n", number2)

	name := "Purva Thakkar" // short variable declaration
	fmt.Printf("%T\n", name)

}

func PrintsDefaultValue() {

	var number uint // implicit type
	var bl bool     // explicit type
	var name string

	fmt.Println("String: ", name) // default value of string is empty string
	fmt.Println("int: ", number)  // default value of int is 0
	fmt.Println("boolean: ", bl)  // default value of bool is false
}

func PrintsConsolAndFmt() {

	fmt.Printf("Hello %T %v \n", 10, 10) // print type of variable

	fmt.Printf("Hello %t \n", true)

	name := "Purva Thakkar"
	fmt.Printf("Hi my name is %s \n", name) // print without quotes
	fmt.Printf("Hi my name is %q \n", name) // print with quotes
	fmt.Printf("Hi my name is %19q \n", name)

}
