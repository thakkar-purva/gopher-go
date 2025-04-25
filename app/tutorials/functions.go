package tutorials

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Function to demonstrate a simple function

func WorkWithFunction() {

	fmt.Println("Hello, World!")
	fmt.Println("This is a simple function in Go.")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter value x :")
	scanner.Scan()
	xValue, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Println("Enter value y :")
	scanner.Scan()
	yValue, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	add(int(xValue), int(yValue)) // function call

	subtract(int(xValue), int(yValue)) // function call

	var ans = returnvalue(int(xValue), int(yValue)) // function call
	fmt.Printf("Return value for Addition of two numbers is %d\n", ans)

	ans1, ans2 := returnAllValues(int(xValue), int(yValue))
	fmt.Printf("Return all values fior functions of two numbers is %d %d\n", ans1, ans2) // function call

	deferFunction() // function call
	fmt.Println("Defer function called")
}

func add(x int, y int) {
	fmt.Println("Addition of two numbers is ", x+y)
}

func returnvalue(x int, y int) int {
	return x + y
}

// this is different way to declare function
func subtract(x, y int) {
	fmt.Println("Subtraction of two numbers is ", x-y)
}

func returnAllValues(x, y int) (int, int) {
	return x + y, x - y
}

func deferFunction() {
	// defer is used to delay the execution of a function until the surrounding function returns
	// it is used to ensure that a function call is performed later in a program's execution
	// it is often used for cleanup tasks, such as closing files or releasing resources
	defer fmt.Println("world")
	defer fmt.Println("Hello")
}
