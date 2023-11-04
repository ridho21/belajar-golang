package main

import "fmt"

func main() {
	fmt.Println("Tipe Data")
	fmt.Println("Numerik")

	var positiveNumber uint8 = 90
	var negativeNumber int = -180
	fmt.Printf("Bilangan Positif : %d\n", positiveNumber)
	fmt.Printf("Bilangan Negative : %d\n", negativeNumber)

	var decimalNumber = 3.1
	fmt.Printf("Bilangan pecahan : %.2f", decimalNumber)

	var message = `Nama saya 'Ridho',
	salam kenal,
	Mari belajar golang`

	fmt.Printf(message)
}
