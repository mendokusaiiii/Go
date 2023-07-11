package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type your birthyear: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at the end of 2023", 2023-input)
}
