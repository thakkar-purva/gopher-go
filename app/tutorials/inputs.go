package tutorials

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// The above code uses the bufio package to read input from the console.
// It creates a new scanner that reads from os.Stdin (the standard input stream).
func PrintsInput() {
	fmt.Println("What is your name?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Hello", name)

	// different way to read integer input.
	fmt.Printf("type the year you were born in : ")
	scanner.Scan()
	year, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You were born in %d \n", year)
	fmt.Println("Thank you !!")
}
