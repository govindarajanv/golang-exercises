package main

import "fmt"
import "runtime"

func main() {
	os := runtime.GOOS

	switch os {
	case "linux":
		fmt.Println("This runs on Linux")
	case "darwin":
		fmt.Println("This runs on Darwin")
	default:
		fmt.Printf("This runs on %s\n",os)
	}
}