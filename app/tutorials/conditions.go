package tutorials

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// public method
func WorkWithConditions() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a value for x : ")
	scanner.Scan()
	x := scanner.Text()
	xVal, _ := strconv.ParseInt(x, 10, 64)

	fmt.Println("Enter a value for y : ")
	scanner.Scan()
	y := scanner.Text()
	yVal, _ := strconv.ParseInt(y, 10, 64)

	isHigher := xVal > yVal
	fmt.Printf("Is %s higher than %s? %t \n", x, y, isHigher)

	andOperator(xVal, yVal) // call private method
	orOperator(xVal, yVal)  // call private method

}

// private method
func andOperator(x int64, y int64) {
	isSame := x == y
	isHigher := x > y
	value := isSame && isHigher // if both are true then only it will be true
	fmt.Printf("and Operator : %t \n", value)
}

// private method
func orOperator(x int64, y int64) {
	isSame := x == y
	isHigher := x > y
	value := isSame || isHigher // if one of is true then only it will be true
	fmt.Printf("or Operator : %t \n", value)
}
