package main

import "fmt"

func main() {
	num := []int{3, 2, 3, 6}
	_, besar := minMax(num)
	// fmt.Println("Min : ", kecil)
	fmt.Println("Max : ", besar)
}

//normal function with return
// func minMax(numbers []int) (int, int) {
// 	min := numbers[0]
// 	max := numbers[0]

// 	for _, n := range numbers {
// 		if n < min {
// 			min = n
// 		}
// 		if n > max {
// 			max = n
// 		}
// 	}
// 	return min, max
// }

//named return value
func minMax(numbers []int) (min int, max int) {
	min = numbers[0]
	max = numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}
