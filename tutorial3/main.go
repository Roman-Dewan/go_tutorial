package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating for our Pizza: ")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

	numRating, error := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if error != nil {
		fmt.Println("error: ", error)
	} else {
		fmt.Println("Thanks for rating", numRating)
		fmt.Printf("Types: %T", numRating)
	}
	// fmt.Printf("Types: %T", input)

}
