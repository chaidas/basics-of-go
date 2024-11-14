package main

import "fmt"

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
	// no need to import functions.go, it's part of same package and thus printData() is available here
	printData()

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
strings
	`
	fmt.Println(str1, str2)

}
