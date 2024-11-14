package main

var name = "Example of global var"

// Example of global var (technically "package" var)
var url = "https://frontendmasters.com"

func main() {
	// function-scoped variables
	message := "Hello from the go_input_output module!"
	price := 34.4
	var age uint8 = 255
	println(message, price, url, age)
	// no need to import functions.go, it's part of same package and thus printData() is available here
	printData()
}
