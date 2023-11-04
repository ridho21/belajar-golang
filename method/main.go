package main

import "fmt"

type triangle struct {
	alas   float64
	tinggi float64
}

type counter struct {
	value int
}

//value receiver -> (t triangle)
func (t triangle) area() float64 {
	return 0.5 * t.alas * t.tinggi
}

//pointer receiver
func (t *triangle) increaseSize() {
	t.alas += 1
	t.tinggi += 1
	fmt.Println(t.alas)
	fmt.Println(t.tinggi)
}

func (c counter) Increment() {
	c.value++
}

func main() {
	instanceTriangle := triangle{10, 12}
	counter := &counter{value: 3}
	counter.Increment()
	area := instanceTriangle.area()
	fmt.Println(area)
	fmt.Println("=============================")
	fmt.Println("Instance Triangle : ", instanceTriangle)
	instanceTriangle.increaseSize()
	fmt.Println("Instance Triangle : ", instanceTriangle)
	fmt.Println(counter)
}
