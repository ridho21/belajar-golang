package main

import "fmt"

type Shape interface {
	getArea() float64
	getPerimeter() float64
}

type Rectangle struct {
	width  float64
	length float64
}

func (r Rectangle) getArea() float64 {
	return r.width * r.length

}

func (r Rectangle) getPerimeter() float64 {
	return 2 * (r.width + r.length)
}

type Square struct {
	side float64
}

func (s Square) getArea() float64 {
	return s.side * s.side
}

func (s Square) getPerimeter() float64 {
	return 4 * s.side
}

// func getAreaOfRectangle(r Rectangle) {
// 	fmt.Println("Luas : ", r.getArea())
// }

// func getAreaOfSquare(s Square) {
// 	fmt.Println("Luas : ", s.getArea())
// }

func getArea(s Shape) {
	fmt.Println("Luas : ", s.getArea())
}

func main() {
	var shape1 Shape = Square{side: 5}
	var shape2 Shape = Rectangle{width: 5, length: 3}
	var shape3 Shape = Rectangle{width: 7, length: 9}
	fmt.Printf("shape1: %#v \n", shape1)
	fmt.Printf("shape2: %#v \n", shape2)
	fmt.Printf("shape3: %#v \n", shape3)
	fmt.Println(shape1, shape2, shape3)
	fmt.Println("=======================================")

	shapes := []Shape{shape1, shape2, shape3}

	for _, shape := range shapes {
		fmt.Printf("%#v memiliki luas %.1f dan keliling %.1f\n", shape, shape.getArea(), shape.getPerimeter())
	}

	fmt.Println("=======================================")
	// getAreaOfRectangle(Rectangle{width: 3, length: 7})
	// getAreaOfSquare(Square{side: 9})

	getArea(Rectangle{width: 2, length: 3})
	getArea(Square{side: 7})

	//Downcasting
	fmt.Println("=======================================")

	var square1 Shape = Square{side: 6}
	fmt.Println("getArea : ", square1.getArea())
	fmt.Println("getPerimeter:", square1.getPerimeter())
	// fmt.Println("side : ", square1.side)

	var originalSquare Square = square1.(Square)
	fmt.Println("getArea : ", originalSquare.getArea())
	fmt.Println("getPerimeter : ", originalSquare.getPerimeter())
	fmt.Println("side: ", originalSquare.side)

	fmt.Println("=======================================")

	var anything any

	anything = "Kenji"
	fmt.Printf("Anything %T: %v \n", anything, anything)

	anything = 20
	fmt.Printf("Anything %T: %v \n", anything, anything)

	anything = false
	fmt.Printf("Anything %T: %v \n", anything, anything)

	anything = []string{"Golang", "Java", "Python"}
	fmt.Printf("Anything %T: %v \n", anything, anything)

}
