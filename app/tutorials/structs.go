package tutorials

import (
	"fmt"
)

// Structs are user-defined types that allow you to group related data together.
// They are similar to classes in other programming languages but are simpler and more lightweight.
// Structs are used to create complex data types that can hold multiple values of different types.

// Defining a struct
type Person struct {
	name string
	age  int
}

func WorkWithStructs() {
	var purva Person = Person{"Purva", 25}
	var milan Person = Person{name: "Milan"}
	fmt.Println(purva.name, purva.age)
	fmt.Println(milan.name, milan.age)

}
