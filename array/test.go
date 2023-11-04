package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func test() {
	// Tulis kode disini
	var arrLength int
	var inputElement string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arrLength, _ = strconv.Atoi(scanner.Text())

	var arr = make([]string, arrLength)

	scanner.Scan()
	inputElement = scanner.Text()
	sArr := strings.Fields(inputElement)
	for i := 0; i < arrLength; i++ {
		arr[i] = sArr[i]
	}
	fmt.Printf("%q", arr)

	for i := range arr {
		arrInt, _ := strconv.Atoi(arr[i])
		if arrInt%2 == 0 {
			fmt.Printf("%s\n", arr[i])
		}
	}
}
