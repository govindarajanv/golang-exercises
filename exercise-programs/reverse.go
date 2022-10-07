package main

import (
		"fmt"
)

func main() {

	var seqLength int
	fmt.Printf("Enter the length of the sequence")
	fmt.Scanln(&seqLength)

	for i:=1; i <= seqLength; i++ {
		defer fmt.Println(i)
	}
}