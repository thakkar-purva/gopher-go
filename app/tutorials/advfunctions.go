package tutorials

import (
	"fmt"
)

func test() {
	fmt.Println("Hello !! ")
}

func PrintsAdvFunction() {
	x := test
	x()
}
