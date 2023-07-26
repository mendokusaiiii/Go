package main

import "fmt"

func inverted(num int) int {
	return num * -1
}

func invertWPointers(num *int) {
	*num = *num * -1
}

func main() {
	num := 20
	invertedNum := inverted(num)
	fmt.Println(invertedNum)

	newNum := 40
	fmt.Println(newNum)
	invertWPointers(&newNum)
}