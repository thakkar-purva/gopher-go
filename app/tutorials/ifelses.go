package tutorials

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WorkWithIfAndElses() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a value for name : ")
	scanner.Scan()
	name := scanner.Text()

	ifConditions(name)
	elseConditions(name)
	elseifConditions(name)

	fmt.Println("How old are you ? ")
	scanner.Scan()
	age := scanner.Text()
	ageValue, _ := strconv.ParseInt(age, 10, 64)

	canRide(ageValue)
}

func ifConditions(name string) {
	if name == "Purva" {
		fmt.Printf("condtion is ? (name == %s) : %t\n", name, name == "Purva")
	}
	fmt.Printf("condtion is ?  (else == %s): \n", name)
}

func elseConditions(name string) {
	if name == "Purva" {
		fmt.Printf("condtion is ?  (name == %s) : %t\n", name, name == "Purva")
	} else {
		fmt.Printf("condtion is ?  (else = %s) : %t\n", name, name != "Purva")
	}
}

func elseifConditions(name string) {
	if name == "Purva" {
		fmt.Printf("condtion is ? (name == %s) : %t\n", name, name == "Purva")
	} else if name == "purva" {
		fmt.Printf("condtion is ? : (else if == %s) %t\n", name, name == "purva")
	} else {
		fmt.Printf("condtion is ? : (else = %s) : \n", name)
	}
}

func canRide(age int64) {
	if age > 18 {
		fmt.Println("Yes you can drive")
	} else {
		fmt.Println("You can not drive")
	}
}
