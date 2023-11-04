package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := devide(5, 1)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Hasil pembagian: ", result)
}

func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Pembagian dengan angka 0")
	}

	return a / b, nil
}
