package main

import (
		"fmt"
		"bufio"
		"os"
		"strconv"
		"strings"
)

func main() {
	fmt.Println("Welcome to my Pizza app")
	fmt.Println("Please rate our Pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	// Read documentation like 'go doc bufio ReadString'
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for your rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to your rating:", numRating + 1)
	}

	
}