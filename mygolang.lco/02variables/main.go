package main

import "fmt"


// Please note capital L which signifies it is a public variable
const LoginToken string = "DUMMY-TOKEN"

func main(){
	var username string = "Govind"
	fmt.Println(username)
	fmt.Printf("Variable of type: %T\n",username)
	
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable of type: %T\n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable of type: %T\n",smallVal)

	// float64 gives more precision
	var smallFloat float32 = 255.45544123456778
	fmt.Println(smallFloat)
	fmt.Printf("Variable of type: %T\n",smallFloat)

	//default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable of type: %T\n", anotherVariable)

	//Implicit type
	var website = "dummy-website"
	fmt.Printf("variable of type: %T\n", website)
	// This will not be supported - change of type for implicit variables like website = 3

	//no var style but allowed only inside a method not as global , := walrus operator
	numberOfUser := 30000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable of type: %T\n", numberOfUser)

	fmt.Println("My login token", LoginToken)
	fmt.Printf("Variable of type: %T\n", LoginToken)

}