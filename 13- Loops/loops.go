package main

import (
	"fmt"
	"time"
)

	func main() {
		fmt.Println("Loops")

		i := 0

		for i > 10 {
			i++
			fmt.Println("Increment")
			time.Sleep(time.Second)
		}
		fmt.Println(i)

		nomes := [3]string{"joao", "pedro", "jose"}

		for _, nome := range nomes {
			fmt.Println(nome)
		}

		for index, letter := range "WORD" {
			fmt.Println(index, string(letter))
		}
	}