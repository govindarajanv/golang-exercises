package main

import "fmt"

type Aquatic interface {
	swim()
}

type Bird struct {
	name string

}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (b Bird) swim() {
	fmt.Println(b.name, "can swim")
}

func main() {
	var bird Aquatic = Bird{"duck"}
	bird.swim()
}