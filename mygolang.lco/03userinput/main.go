package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println ("Enter the rating:")
	//comma, ok syntax || error ok
	input, _ := reader.ReadString('\n')  // input, err := similar to try catch
	fmt.Println("Thanks for your rating",input)
	fmt.Printf("Type of the rating is %T\n", input)
}