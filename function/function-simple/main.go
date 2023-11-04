package main

import "fmt"

// var name = "Ridho"

func main() {
	helloworld()
	greetWorld("Ridho")
}

func helloworld() {
	fmt.Println("Hello World")
}

func greetWorld(name string) {
	fmt.Println("Hello", name)
}
