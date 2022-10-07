// https://go.dev/tour/moretypes/6
package main

import (
	"fmt"
)

type Rectangle struct {
	length int
	breadth int
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

func (r *Rectangle) swap() {
	var t int
	t = r.length
	r.length = r.breadth
	r.breadth = t
}

func main() {
	var l,b int
	fmt.Printf("Enter the length")
	fmt.Scanln(&l)
	fmt.Printf("Enter the breadth")
	fmt.Scanln(&b)

	r := Rectangle{l,b}
	r2 := Rectangle{length: 8}
	r3 := &Rectangle{}
	r4 := &Rectangle{7,6}

	fmt.Println(r)
	fmt.Println(r2)
	fmt.Println(*r3)
	fmt.Println(r3.length)

	// Print areas
	fmt.Println("Areas of the rectangles")
	fmt.Println(r.area())
	fmt.Println(r2.area())
	fmt.Println(r3.area())
	fmt.Println(r4.area())

	// Swapping

	fmt.Println(r)
	r.swap()
	fmt.Println(r)

}