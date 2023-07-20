package main

import "fmt"

func mathCalculus(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	sum, sub := mathCalculus(20, 10)
	fmt.Println(sum, sub)
}