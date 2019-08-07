package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banan", "melon"}
	var afruits = fruits[0:2]

	fmt.Println(fruits)
	fmt.Println(afruits)
	afruits[1] = "anu"
	fmt.Println(fruits)
	fmt.Println(afruits)

}
