package main

import "fmt"

type person struct {
	name		string
	lastname	string
	age			uint8
	height		uint8
}

type student struct {
	person
	course	string
	collage	string
}

func main() {
	fmt.Println("heranca")

	p1 := person{"Victor", "Cabral", 25, 186}
	fmt.Println(p1)

	e1 := student{p1, "Software Engeneering", "PUCPR"}
	fmt.Println(e1.age)
}