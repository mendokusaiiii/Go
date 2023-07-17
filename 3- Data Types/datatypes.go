package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int8 = 1
	fmt.Println(number)

	var number2 uint32 = 1000
	fmt.Println(number2)

	// alias
	// INT32 = RUNE

	var number3 rune = 12456
	fmt.Println(number3)

	// BYTE = INT8
	var number4 byte = 123
	fmt.Println(number4)

	var floatNUmber float32 = 123.4
	fmt.Println(floatNUmber)

	floatNUmber2 := 1235.66666
	fmt.Println(floatNUmber2)

	var str string = "oi"
	fmt.Println(str)

	str2 := "O"
	fmt.Println(str2)

	var booleand bool = true
	fmt.Println(booleand)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
