package main

import (
	"fmt"
	"gopher-go/app/tutorials" // Import the tutorials package
)

// entry point for go project  NOTE: has to be main() method in main package
// file name can be anything.
func main() {
	fmt.Println("Hello World !!")
	tutorials.PrintsTutorial() // Call the exported function from the tutorials package
}
