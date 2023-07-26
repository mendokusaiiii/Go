package main

import "fmt"

func closure() func() {
	txt := "Inside closure"

	function := func() {
		fmt.Println(txt)
	}

	return function
}

func main() {
	txt := "Inside main func"
	fmt.Println(txt)

	newFunc := closure()
	newFunc()
}