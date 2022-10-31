package main

import (
	"fmt"
)

func main(){

	var myInt int
	var myIntPointer *int

	myInt = 5
	myIntPointer = &myInt

	fmt.Println("Value at myIntPointer is", *myIntPointer)
}
