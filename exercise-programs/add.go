package main

import "fmt"

func add(x, y int) int {
	return x+y
}

func main() {
	var number1, number2 int
	fmt.Printf("Enter the first number:")
	fmt.Scanln(&number1)
	fmt.Printf("Enter the second number:")
	fmt.Scanln(&number2)

	
	fmt.Println("Sum is", add(number1,number2))
}