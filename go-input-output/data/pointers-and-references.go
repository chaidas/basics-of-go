package data

import "fmt"

func PointersAndRefs() {

	// Defer example
	defer fmt.Println("Ok bye friends!")

	fmt.Println("--- Pointers & Refs ---")
	age := 46
	fmt.Println("My age is: ", age)

	// Declare a pointer to the age integer
	agePointer := &age
	fmt.Println("Memory address of age:", agePointer)

	// Dereferencing the pointer to get the value
	fmt.Println("Value of age (through the pointer):", *agePointer)

	// Modifying the value of the variable _through_ the pointer
	*agePointer = 47
	fmt.Println("Value of age (after modification):", agePointer)

	// Playing around with Printf (%v stands for "value")
	fmt.Printf("The pointer is %v and the value is %v\n\n\n", agePointer, *agePointer)

	if age > 44 {
		panic("OH SHITE I'M NOW OVER 44")
	}

}
