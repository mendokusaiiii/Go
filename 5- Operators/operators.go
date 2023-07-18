package main

import "fmt"

func main() {
	// Aritmeticos
	sum := 1 + 2
	sub := 1 -2
	div := 1 / 2
	mul := 1 * 2
	resdiv := 1 % 2

	fmt.Println(sum, sub, div, mul, resdiv)
	// Voce nao pode fazer nada no Go que são de tipos diferentes
	// Ex: Somar um int8 com um int16

	// Atribuicao
	var variable1 string = "Hey"
	variable2 := "Hey you"
	fmt.Println(variable1, variable2)

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores Logicos
	vtrue, vfalse := true, false
	fmt.Println(vtrue && vfalse)
	fmt.Println(vtrue || vfalse)
	fmt.Println(!vtrue)

	// Operadores Unarios
	number := 10
	number++
	number += 20
	number--
	number -= 10
	number *= 2
	number /= 2
	number %= 3
	fmt.Println(number)

	// NAO EXISTE OPERADOR TERNARIO NO GO !
}