package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	// Isi struct ini
	name string
	age  int
}

func (s *Student) Introduction() {
	// Tulis kode disini
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s.name = scanner.Text()
	scanner.Scan()
	s.age, _ = strconv.Atoi(scanner.Text())
	fmt.Print(s)
	fmt.Printf("Hello, my name is %s. Im %d years old", s.name, s.age)
}

func main() {
	// Tulis kode disini
	var student1 = Student{}
	student1.Introduction()
	fmt.Println(student1)
}
