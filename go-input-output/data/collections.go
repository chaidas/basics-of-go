package data

import "fmt"

// init() method is executed typically when you start your app
// A function where you can initialise things
// You can have more than one init() function in the same package in the same file!
// (Careful: init() is _not_ main)
func init() {

	var countries [10]string
	countries[0] = "Argentina"
	countries[1] = "Brazil"
	countries[8] = "USA"

	// Getting the length of arrays
	fmt.Println("Length of Countries array", len(countries))

	// Slice literal example
	countriesSlice := []string{"France", "Japan", "Brazil"}
	fmt.Println(countriesSlice)

	// Map literal example
	wellKnownPorts := map[string]int{
		"http":  80,
		"https": 443,
	}
	fmt.Println(wellKnownPorts)

}
