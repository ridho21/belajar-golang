package main

import "fmt"

func main() {
	numberSlice := []int{2, 7, 9, 4}
	fmt.Println("Panjang Slice: ", len(numberSlice))
	fmt.Println("Kapasitas Slice : ", cap(numberSlice))

	scores := make([]int, 3, 7)
	scores[0] = 23
	scores[1] = 33
	scores[2] = 43

	fmt.Println(scores)

	scores2 := make([]int, 4)
	scores2[0] = 23
	scores2[1] = 33
	scores2[2] = 43
	scores2[3] = 53

	fmt.Println(len(scores2))
	fmt.Println(cap(scores2))

	//perbedaan array dan slice

	//array
	buah := [4]string{"Mangga", "Jeruk", "Durian", "Markisa"}
	fmt.Println(buah)

	//slice
	villain := make([]string, 3, 5)
	villain[0] = "Madara"
	villain[1] = "Obito"
	villain[2] = "Pain"
	fmt.Println("Panjang villain : ", len(villain))
	fmt.Println("Kapasitas villain : ", cap(villain))
	villain = append(villain, "Itachi")
	fmt.Println(villain)

	fmt.Println("========================================")
	//pass by value
	var numbers = [4]int{2, 1, 6, 8}
	var anotherNumbers = numbers
	fmt.Println("Numbers: ", numbers)
	fmt.Println("Another Numbers :", anotherNumbers)
	anotherNumbers[1] = 11
	fmt.Println("Numbers: ", numbers)
	fmt.Println("Another Numbers :", anotherNumbers)

	fmt.Println("========================================")
	//pass by reference
	var prices = []int{20000, 13000, 10000}
	var anotherPrices = prices
	fmt.Println("Prices: ", prices)
	fmt.Println("Another Price :", anotherPrices)
	anotherPrices[1] = 15000
	fmt.Println("Prices: ", prices)
	fmt.Println("Another Price :", anotherPrices)

}
