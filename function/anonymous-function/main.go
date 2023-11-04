package main

import "fmt"

func main() {
	//Anonymous function
	func() {
		fmt.Println("Hello Wordl!")
	}()

	//Anonymous function to variable
	halo := func() {
		fmt.Println("Halo Dunia!!")
	}

	halo()

	//Passing argument to anonymous func
	func(word string) {
		fmt.Println(word)
	}("Enigma Camps")

	//Passing argument dalam variable func
	hello := func(name string) {
		fmt.Println("Hello", name)
	}
	hello("Ridho")

	//Passing anonymous function sebagai argumen
	greetEnglish := func(name string) string {
		return "Hola " + name
	}

	greetRussian := func(name string) string {
		return "Blyat " + name

	}
	greet("Ridho", greetEnglish)
	greet("Ridho", greetRussian)

	add := func(num1 int, num2 int) int {
		return num1 + num2
	}
	multiply := func(num1 int, num2 int) int {
		return num1 * num2
	}
	calculate(2, 2, add)
	calculate(2, 4, multiply)

}

func greet(name string, f func(name string) string) {
	fmt.Println(f(name))
}

func calculate(a int, b int, operator func(x int, y int) int) {
	fmt.Println(operator(a, b))
}
