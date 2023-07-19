package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string {
		"name": "Victor",
		"lastname": "Cabral",
	}

	fmt.Println(user)

	// anihamento de maps
	user2 := map[string]map[string]string {
		"name": {
			"firstname": "Joao",
			"lastname": "Silva",
		},
		"address": {
			"addre": "Rua Ficticia",
			"number": "456",
		},
	}

	fmt.Println(user2)
	delete(user2, "address")
}