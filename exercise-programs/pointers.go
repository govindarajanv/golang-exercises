package main

import ( 
		"fmt"
		)


func main() { 
	var a,b int
	fmt.Printf("First number:")
	fmt.Scanln(&a)
	fmt.Printf("Second number:")
	fmt.Scanln(&b)
	fmt.Println("Before swapping",a,b)
	swap (&a,&b)
	fmt.Println("After swapping", a,b)
}

func swap(x,y *int) {
	var t int
	t = *x
	*x = *y
	*y = t

}