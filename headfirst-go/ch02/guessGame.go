// implement as show in pg#146
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	//"reflect"
	"strconv"
	"strings"
	"time"
)

func main(){
	// get the current time and date as integers
	seconds := time.Now().Unix()
	//fmt.Println("type of second:",reflect.TypeOf(seconds))

	// Seed tjhe random number generator
	rand.Seed(seconds)

	// Generate an integer between 1 and 100
	target := rand.Intn(100)
	fmt.Println("A number is chosen between 1 and 100. Can you guess it?")

	// Read user's input from the keyboard
	reader := bufio.NewReader(os.Stdin)

	//success flag initialized to false
	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("\nYou have ",10-guesses, " left")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess,err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Your guess was low")
		} else if guess > target {
			fmt.Println("Your guess was high")
		} else {
			fmt.Println("You guessed it right")
			success = true
			break
		}

	}
	if !success {
		fmt.Println("You lost.Chosen number was ",target)
	}

}

