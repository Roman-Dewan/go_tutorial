package main

import (
	"fmt"
)

func main() {
	var username string = "roman"
	fmt.Println(username)

	var result int = add(5, 2)
	fmt.Println(result)

	var number uint8 = 216
	fmt.Println(number)

	a := 20
	fmt.Printf("a : %d", a)
}

func add(a int, b int) int {
	var sum int = a + b
	return sum
}
