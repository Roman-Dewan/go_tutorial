package main

import "fmt"

func main() {
	var alphabet = []string{"A", "B", "C", "D", "E"}
	fmt.Println("fruit list: ", alphabet)
	fmt.Printf("type of : %T\n", alphabet) // Added \n for formatting

	// Reslice to drop the first element
	alphabet = alphabet[1:]
	fmt.Println("After slicing: ", alphabet) // Output: [B C D E]

	// remove element from slices
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
