package main

import "fmt"

func main() {
	fmt.Println(add(2, 3))
}

func add(a int, b int) int {
	result := a + b
	return result
}
