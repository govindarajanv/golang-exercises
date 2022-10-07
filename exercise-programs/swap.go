package main

import "fmt"

func swap(a, b string) (string, string) {
	return b,a
}

func main(){
	var text1, text2 string
	fmt.Printf("Enter your first text:")
	fmt.Scanln(&text1)
	fmt.Printf("Enter your second text:")
	fmt.Scanln(&text2)

	a,b := swap(text1,text2)

	fmt.Println("Swapped output is ", a,b)

}