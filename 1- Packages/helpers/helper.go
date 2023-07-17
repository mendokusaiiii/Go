package helpers

import "fmt"

// Go não é orientado a objetos, então para poder usar funções de outros pacotes diferenciamos com letras maiculas e minuculas
func Writings() {
	fmt.Println("Typing from a helper package")
	writings2()
}
