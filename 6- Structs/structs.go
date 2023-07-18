package main

import "fmt"

type user struct {
	name string 
	age uint8
	address address
}

type address struct {
	zipcode string
	number uint8
}

func main() {
	fmt.Println("Structs")

	// Primeira maneira de trabalhar com struct
	var u user
	u.name = "Victor"
	u.age = 25
	fmt.Println(u)

	// Segunda maneira
	aaddress := address{"Rua ficticia", 43}
	u2 := user{"Victor", 25, aaddress}
	fmt.Println(u2)

	// Terceira maneira
	u3 := user{name: "Victor"} // Esse metodo é quando somente tenho um dos metodos
	fmt.Println(u3)
}