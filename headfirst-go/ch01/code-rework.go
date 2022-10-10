package main

import (
	"fmt"
)

func main() {
	var pebbleweight,rockweight,boulderweight = 0.1, 1.2, 502.4
	total_weight := pebbleweight + rockweight + boulderweight
	fmt.Println(total_weight)
}