package main

import (
	"access-modifier/library"
	"fmt"
)

// exporter = huruf awal besar
// importer = huruf awal kecil
func main() {
	fmt.Println("HourInADay : ", library.HourInADay)
	fmt.Println("Name : ", library.Name)

	fmt.Println(library.DaysToMinutes(3))
	// fmt.Println(library.daysToHour(5))
}
