package main

import (
	"fmt"
	"modules/helpers"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hey")
	helpers.Writings()
	erro := checkmail.ValidateFormat("victordev@gmail.com")
	fmt.Println(erro)
}
