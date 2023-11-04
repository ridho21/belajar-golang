package main

import "fmt"

func main() {
	var name = "ridho"
	fmt.Println("name : ", name)
	fmt.Println("alamat memori : ", &name)

	var age = 25
	fmt.Println("age : ", age)
	fmt.Println("alamat memori", &age)
}
