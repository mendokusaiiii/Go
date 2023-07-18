package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// no Go voce pode ter mais de um retorno em uma função. Exemplo abaixo

func matCalculus(n1, n2 int8) (int8, int8) {
	sumMethod := n1 + n2
	subMethod := n1 - n2
	return sumMethod, subMethod
}

func main() {
	sums := sum(10, 20)
	fmt.Println(sums)

	var f = func(txt string) string {
		fmt.Println("Function f")
		return txt
	}

	result := f("Test")
	fmt.Println(result)

	sumResult, subResult := matCalculus(20, 30) // Se por algum motivo voce não quiser utilizar uma das variaveis, voce pode simplesmente utilizar um "_"
	fmt.Println(sumResult, subResult) 
}
