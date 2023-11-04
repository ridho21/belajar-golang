package main

import "fmt"

func main() {
	total := sum(3, 4, 4, 2)
	fmt.Println(total)
}

func sum(numbers ...int) int {
	// fmt.Println(numbers[0])
	// fmt.Println(numbers[1])
	// fmt.Println(numbers[2])
	// fmt.Printf("%T \n", numbers)

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
