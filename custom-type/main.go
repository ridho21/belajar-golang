package main

import "fmt"

type Patient struct {
	Name string
	Age  int
	Celcius
}

type Celcius float64

func (c Celcius) toFarenheit() float64 {
	return float64(c*9/5 + 32)
}

func (c *Celcius) add(value float64) {
	*c += Celcius(value)
}

func main() {
	var temperature Celcius = 20.0
	fmt.Println(temperature)
	fmt.Println("Suhu di ruangan ini : ", temperature.toFarenheit())
	temperature.add(3)
	fmt.Println(temperature)
	fmt.Println("================================================================")
	newPatient := Patient{Name: "ridho", Age: 25, Celcius: 39}
	fmt.Printf("New Patient : %+v", newPatient)
}
