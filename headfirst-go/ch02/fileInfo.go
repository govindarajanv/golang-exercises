package main

import (
	"fmt"
	"os"
	"log"
)

func main() { 
	fileInfo, err := os.Stat("README.md")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file with name %s is %d\n", fileInfo.Name(),fileInfo.Size())
}
