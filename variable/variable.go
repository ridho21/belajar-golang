package main

import (
	"fmt"
)

func main() {
	var firstname string = "Ridho"

	var lastname string = "Mahendra"

	fmt.Println(firstname, lastname)
	fmt.Printf("Halo, Ridho Mahendra")
	fmt.Printf("\nHalo, %s %s", firstname, lastname)

	var (
		age     int
		address string
	)

	age = 25
	address = "Jakarta"

	fmt.Printf("\nHalo, saya %s %s berumur %d tinggal di %s", firstname, lastname, age, address)

	var (
		bootcampName, bootcampAddress = "Enigma", "Jakarta"
	)

	fmt.Printf("\nBootcamp %s %s", bootcampName, bootcampAddress)

	day := "Monday"
	date := "31"
	month := "October"

	fmt.Println(day, date, month+"2023")

	const phi = 3.14

	fmt.Println(phi)

	bootcamp, _ := "Enigma Camp", "Aktif7"

	fmt.Println(bootcamp)
}
