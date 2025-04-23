package tutorials

import (
	"fmt"
)

func WorkWithArrays() {
	fmt.Println("All about arrays")
	/* arrays are fixed size
	arrays are value types
	arrays are not resizable
	arrays are not reference types */

	var arr [5]int
	// by default all values are initialized to 0
	fmt.Println(arr) // [0 0 0 0 0]

	arr[0] = 1
	fmt.Println(arr) // [1 0 0 0 0]

}
