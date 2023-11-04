package main

import "fmt"

type kendaraan struct {
	merek string
	tahun int
	model string
	harga int
}

func main() {
	var a kendaraan

	a.merek = "Toyota"
	a.tahun = 2009
	a.model = "Rush"
	a.harga = 200000000

	fmt.Println(a)

	fmt.Println("--------------------------------")

	var b = kendaraan{}
	b.merek = "Wuling"
	b.tahun = 2023
	b.model = "REV"
	b.harga = 400000000
	fmt.Println("b :", b)

	var c = kendaraan{"Honda", 2023, "Brio", 250000000}
	fmt.Println("c :", c)

	fmt.Println("--------------------------------")

	var d = kendaraan{model: "Xenia", merek: "Toyota"}
	fmt.Println("d : ", d)

	fmt.Println("--------------------------------")
	var c2 *kendaraan = &c
	c2.model = "awea"
	fmt.Println(c)
	fmt.Println("c2 : ", c2)

	fmt.Println("--------------------------------")
	var k = new(kendaraan)
	fmt.Println(&k)
}
