package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) +fibonacci(position-1) 
	/* 
	Entendendo um pouco melhor de como funciona esse ultimo return
	Esse primeiro fibonacci(position-2) está pegando a posição que foi passada(a que queremos saber)
	e esta diminuindo duas casas que a antecedem. Consequentemente a fibonacci(position-1) esta pegando uma
	casa antes da que foi passada.
	*/
}

func main() {
	fmt.Println()

	position := uint(10)
	fmt.Println(fibonacci(position))
}