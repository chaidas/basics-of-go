package data

import "fmt"

// Countries Array example
var Countries [10]string

// Slice An example of a slice
var Slice []int

// Codes Map example
var Codes map[int]string

// init() method is executed typically when you start your app
// A function where you can initialise things
// You can have more than one init() function in the same package in the same file!
// (Careful: init() is _not_ main)
func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[8] = "USA"

	// Getting the length of arrays
	qty := len(Countries)

	fmt.Println("Length of Countries array", qty)

	// Slice literal example
	countries := []string{"France", "Japan", "Brazil"}
	fmt.Println(countries)

	// Map literal example
	wellKnownPorts := map[string]int{
		"http":  80,
		"https": 443,
	}
	fmt.Println(wellKnownPorts)

}
