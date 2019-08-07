package main

import "fmt"

func findmax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	var getNumbers = func() []int {
		return res
	}
	return len(res), getNumbers
}

func main() {
	var max = 3
	var num = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var panjang, numbernya = findmax(num, max)
	var theNumbers = numbernya()
	fmt.Println("numbers\t:", theNumbers)
	fmt.Printf("find \t: %d\n\n", max)
	fmt.Println("found \t:", panjang)

}
