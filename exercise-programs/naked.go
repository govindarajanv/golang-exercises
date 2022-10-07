package main

import "fmt"

// naked return
func swap(a, b int) (x, y int) {
	x = b
	y = a
	return
}

func main() {
	var num1,num2 int
	fmt.Printf("First number:")
	fmt.Scanln(&num1)
	fmt.Printf("Second number:")
	fmt.Scanln(&num2)

	a,b := swap(num1,num2)

	fmt.Println("swap of two numbers is", a,b )
}