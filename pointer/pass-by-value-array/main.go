package main

import "fmt"

func main() {
	var numbers = [4]int{1, 2, 3, 4}
	var anotherNumbers = numbers

	fmt.Println("number", numbers)
	fmt.Println("another", anotherNumbers)

	numbers[0] = 9
	fmt.Println("number", numbers)
	fmt.Println("another", anotherNumbers)
}
