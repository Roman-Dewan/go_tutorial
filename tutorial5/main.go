package main

import "fmt"

func main() {
	/*
		new() -- allocate memory but no init
			zero storage -> only read/view

		make() -- allocate memory and init
			non zero sorage -> view and physical address.

			Pointers

	*/

	fmt.Println("welcome")

	// var ptr *int
	// fmt.Println("value of pointer ", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of actual ptr is: ", ptr)
	fmt.Println("Value of actual  *ptr is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("old value: ", myNumber)
	fmt.Println("new value: ", *ptr)
	fmt.Println("old value: ", myNumber)

}
