package main

import "fmt"

func function1() {
	fmt.Println("Executin 1")
}

func function2() {
	fmt.Println("Executin 1")
}

func studentApproved(n1, n2 float32) bool {
	defer fmt.Println("Showing result")
	fmt.Println("Student informations")

	avarege := (n1 + n2) / 2

	if avarege >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(studentApproved(10, 7))
	fmt.Println("Defer")
	defer function1()
	function2()
}