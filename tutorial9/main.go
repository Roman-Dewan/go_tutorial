package main

import "fmt"

func main() {
	person := User{Name: "roman", Email: "roman.dewan18@gmail.com", Status: false, Age: 23}
	fmt.Println(person)
	fmt.Printf("person details are %+v", person)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
