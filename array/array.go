package main

import "fmt"

func main() {
	var fruits = [4]string{
		"Apel", "Pisang", "Semangka"}

	fruits[1] = "test"

	fmt.Println(fruits)

	//deklarasi tanpa mengisi nilai
	var score [3]int

	score[0] = 1
	score[1] = 2
	score[2] = 3
	fmt.Println(score)

	//deklarasi tanpa menentukan panjang array
	var hari = [...]string{"senin", "selasa"}

	fmt.Println(hari)
	fmt.Println("Panjang element : ", len(hari))

	for i := 0; i < len(hari); i++ {
		fmt.Printf("Element ke %d : %s\n", i, hari[i])
	}

	for i, day := range hari {
		fmt.Printf("Element ke %d : %s\n", i, day)
	}

	for _, day := range hari {
		fmt.Printf("Hari : %s\n", day)
	}

	for i := range hari {
		fmt.Printf("Index : %d\n", i)
	}

	for _, day := range hari {
		if day == "senin" {
			fmt.Printf("Index ke %s", day)
		}
	}

	var number = []int{2, 4, 5, 3}
	for i := range number {
		number[i] *= 2
	}
	fmt.Print("\n", number)
}
