package main

import "fmt"

func restoreExecution() {
	if r := recover(); r != nil {
		fmt.Println("Execution restored")
	}
}

func studentApproved(n1, n2 float64) bool {
	defer restoreExecution()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA É EXATAMENTE 6!") // A função panic() vai interromper o fluxo normal do programa e vai parar tudo

}

func main() {
	fmt.Println(studentApproved(6,6))
}