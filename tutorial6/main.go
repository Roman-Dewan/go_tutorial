package main

import "fmt"

func main() {
	var fruitList []string

	fruitList = append(fruitList, "apple")
	fruitList = append(fruitList, "mango")
	fruitList = append(fruitList, "lichi")

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
}
