package main

import (
	"fmt"
)

func main() {
	var x = 4
	var y = x

	y = y + 1
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

	fmt.Println("x: ", &x)
	fmt.Println("y: ", &y)

	var a = 3
	increase(a)
}

func increase(n int) {
	n = n + 1
	fmt.Println("n: ", n)
}
