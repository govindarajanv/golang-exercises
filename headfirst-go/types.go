package main

import ( 
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(2))
	fmt.Println(reflect.TypeOf(1.25))
	fmt.Println(reflect.TypeOf(float64(2)))
}