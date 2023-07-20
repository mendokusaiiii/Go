package main

import "fmt"

func sum(num ...int) int {
	total := 0
	for _, nums := range num {
		total += nums
	} 

	return total
}

func write(txt string, nums ...int) {
	for _, num := range nums {
		fmt.Println(txt, num)
	}
}

func main() {
	sumTotal := sum(20, 25, 67)
	fmt.Println(sumTotal)

	write("Hey", 1, 2)
}