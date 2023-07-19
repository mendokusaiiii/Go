package main

import "fmt"

func weekDays(num int) string {
	switch num {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	default:
		return "Invalid"
	}
}

// outra maneira de usar o switch
func weekDays2(num int) string {
	var weekDay string

	switch {
	case num == 1:
		weekDay = "Sunday"
		fallthrough // ele entra na condição e automaticamente joga pra condição de baixo. Já retornando a proxima condição
	case num == 2:
		weekDay = "Monday"
	default:
		weekDay = "Invalid"
	}
	return weekDay
}

func main() {
	fmt.Println("Switch")
	days := weekDays(1)
	fmt.Println(days)
}