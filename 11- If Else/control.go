package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	number := 10

	if number > 15 {
		fmt.Println("Bigger than 15")
	} else {
		fmt.Println("Lower or equal to 15")
	}

	// if init
	if otherNumber := number; otherNumber > 0 {
		fmt.Println("Bigger than zero")
	} else {
		fmt.Println("Lower than zero")
	}
}