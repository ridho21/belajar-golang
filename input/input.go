package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var firstname string
	var lastname string
	var age int

	fmt.Print("Enter Your Name : ")
	fmt.Scanln(&firstname, &lastname, &age)

	fmt.Println("Your name is : ", firstname, lastname, age)

	fmt.Println("======================================")

	var name string
	var address string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Data diri siswa bootcap")
	fmt.Print("Masukkan Nama anda : ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Print("Masukkan umur anda : ")
	scanner.Scan()
	age, _ = strconv.Atoi(scanner.Text())
	fmt.Print("Masukkan alamat anda : ")
	scanner.Scan()
	address = scanner.Text()

	fmt.Printf("%s | %d | %s", name, age, address)

}
