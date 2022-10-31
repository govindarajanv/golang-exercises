package main

import (
	"fmt"
)

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}

func negate(value *bool) {
	*value = !*value
}
