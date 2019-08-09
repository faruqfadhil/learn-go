package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}
	fmt.Println(allStudents[0].name)
}
