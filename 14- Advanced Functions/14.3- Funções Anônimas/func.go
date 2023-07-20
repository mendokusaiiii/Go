package main

import "fmt"


func main() {
	returns := func (txt string) string  {
		return fmt.Sprintf("Receive -> %s", txt) // SprintF concatena informações
	}("Parameter")

	fmt.Println(returns)
}