package main

import "fmt"

func variadicCoba(numbers ...int) {
	fmt.Println(numbers[0])
}
func main() {
	variadicCoba(1, 2, 3, 4, 5)
}
