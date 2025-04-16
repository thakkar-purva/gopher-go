package tutorials

import (
	"bufio"
	"fmt"
	"os"
)

func PrintsInput() {
	fmt.Println("What is your name?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Hello", name)
}

// The above code uses the bufio package to read input from the console.
// It creates a new scanner that reads from os.Stdin (the standard input stream).
