package main

// Importing fmt so that our PrintLn
import "fmt"

func printData() {
	fmt.Print("Hello ")
	fmt.Println("World")
	// Package var (from main.go) available here given it's a package var and thus in scope
	fmt.Println(name)
}
