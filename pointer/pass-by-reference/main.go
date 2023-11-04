package main

import "fmt"

func main() {
	//& -> ampersan
	//* -> asteris berguna untuk deklarasi variabel pointer

	var x int = 4
	var y *int = &x

	fmt.Println("x", x)
	fmt.Println("alamat x", &x)
	fmt.Println("y", y)
	fmt.Println("alamat y", &y)

	fmt.Println("Nilai dari reference y :", *y)

	var a int = 3
	increase(&a)
	fmt.Println("a : ", a)
}

func increase(n *int) {
	*n = *n + 1
	fmt.Println("n di function increase : ", *n)
}
