package main

import "fmt"
import "go-input-output/data"

var name = "Example of global var"

// Example of global var (technically "package" var)
var url = "https://frontendmasters.com"

func init() {
	fmt.Println("This is from init() hello!")
}

func main() {
	// function-scoped variables
	message := "Hello from the go_input_output module!"
	price := 34.4
	var age uint8 = 255
	fmt.Println(message, price, url, age)

	// "Numbers & collections" chapter
	// id := 4
	// price := 45.4
	// "Mismatched types int and float64)
	// println(id + price)
	// priceAsInt := int(price)
	// idAsFloat := float32(isd)

	str1 := "Single line text"
	str2 := `Better to use backtick 
for multi-line
strings`
	str3 := `<a href="#">Test</a>`
	fmt.Println(str1, str2, str3)

	fmt.Println("---------------------------------")

	// Playing around with collections
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

	fmt.Println("---------------------------------")

	// Functions playground
	// There's no need to import anything since functions.go is of the same package ("main")
	functionPlayground()

	stateTax, cityTax := calculateTax(100)
	fmt.Println(stateTax, cityTax)

	fmt.Println(calculateTaxWithNames(100))

	data.PointersAndRefs()

}

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

// Example with named returns
func calculateTaxWithNames(price float32) (stateTax float32, cityTax float32) {
	// Return values are now available as local vars and can be manipulated...
	stateTax = price * 0.09
	cityTax = price * 0.02
	// ... and we can simply return
	return
}
