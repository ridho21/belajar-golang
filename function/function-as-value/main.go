package main

import "fmt"

func main() {
	var sum func(int, int) int = add
	fmt.Println("Hasil: ", sum(3, 4))
}

func add(a, b int) int {
	return a + b
}
