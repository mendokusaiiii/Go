package main

import "fmt"

func main() {
	fmt.Println("Arrays & Slices")

	var arr1 [5]string
	arr1[0] = "Position 1"
	fmt.Println(arr1)

	array2 := [2]string{"Pedro", "Jõao"}
	fmt.Println(array2)

	arr3 := [...]int{1,2}
	fmt.Println(arr3)

	// SLICE
	slice := []int{10, 20, 30, 40, 50,  60, 70, 80, 90}
	fmt.Println(slice)
	slice = append(slice, 100)
	fmt.Println(slice)
	
	// Slice não é um array! Ele aponta para um array
	// Slice é uma fatia de um array


	// Arrays internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice = append(slice, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	// Quando o slice estoura o seu "limite", o Go dobra o tamanho do seu slice para caberem mais valores nele
	// O slice é uma lista sem tamanho fixo


}