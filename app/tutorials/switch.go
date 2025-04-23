package tutorials

import (
	"fmt"
)

func WorkWithSwitch() {
	fmt.Println("all about switch statements")

	ans := 2

	switch ans {
	case 1:
		fmt.Println("1: One")
	case 2:
		fmt.Println("2: two")
	}

	/* can also put conditions in switch statement, default case is optional*/
	switch {
	case ans > 1:
		fmt.Println("grater than one")

	case ans < 1:
		fmt.Println("less than one")
	default:
		fmt.Println("default case")
	}
}
