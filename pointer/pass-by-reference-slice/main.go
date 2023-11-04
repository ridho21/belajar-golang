package main

import "fmt"

func main() {
	var numbers = []int{4, 3, 7, 11}
	var anotherNumbers = numbers

	fmt.Println("numbers : ", numbers)
	fmt.Println("anotherNumbers : ", anotherNumbers)

	anotherNumbers[1] = 9

	fmt.Println("numbers : ", numbers)
	fmt.Println("anotherNumbers : ", anotherNumbers)

	fmt.Println("=======================================")

	var score = []int{7, 8, 6, 9}
	multiplyBy10(score)
	fmt.Println("score : ", score)
}

func multiplyBy10(numbers []int) {
	for i := range numbers {
		numbers[i] = numbers[i] * 10
	}
	fmt.Println("numbers di function: ", numbers)
}
