package main

import "fmt"

func main() {
	var variable1 int = 10
	var variable2 int = variable1
	fmt.Println(variable1, variable2)
	variable2++
	fmt.Println(variable1, variable2)

	// PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variable3 int
	var pointer *int

	variable3 = 100
	pointer = &variable3

	fmt.Println(variable3, pointer)

	variable3 = 150
	fmt.Println(variable3, *pointer)
}