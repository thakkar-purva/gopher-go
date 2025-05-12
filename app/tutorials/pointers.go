package tutorials

import (
	"fmt"
)

func WorkWithPointers() {
	x := 7
	y := &x
	fmt.Println(x, y)

	*y = 8
	fmt.Println(x, y)

	toChange := "Hello"
	changeValue(&toChange)
	fmt.Println(&toChange)
	changeValue2(toChange)
	fmt.Println(toChange)

	pointer()

}

func changeValue(str *string) {
	*str = "changed"
}

func changeValue2(str string) {
	str = "changed"
}

func pointer() {
	a := "Hello"
	var pointer *string = &a
	fmt.Println(*pointer, &pointer)

}
