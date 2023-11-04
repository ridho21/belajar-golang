package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var bamboo []int
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	//Tulis kode disini
	scanner.Scan()
	amountBamboo, _ := strconv.Atoi(scanner.Text())
	bamboo = addBamboo(amountBamboo)

	scanner.Scan()
	cutAmount, _ := strconv.Atoi(scanner.Text())
	cutBamboo(cutAmount)
}

func addBamboo(amount int) []int {
	sekatBamboo := make([]int, amount)
	for i := range sekatBamboo {
		scanner.Scan()
		inputSekat, _ := strconv.Atoi(scanner.Text())
		sekatBamboo[i] = inputSekat
	}
	return sekatBamboo
}

func cutBamboo(amount int) []int {
	for i := 0; i < amount; i++ {
		fmt.Printf("Potongan ke- %d\n", i+1)
		for j := 0; j < len(bamboo); j++ {
			if bamboo[j] != 0 {
				bamboo[j] = bamboo[j] - 1
			}
			fmt.Println(bamboo[j])
		}
	}
	return bamboo
}
