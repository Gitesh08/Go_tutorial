package main

import "fmt"

func main() {
	var username string = "Gitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isBool bool = true
	fmt.Println(isBool)
	fmt.Printf("Variable is of type %T \n", isBool)

	var Number int = 123
	fmt.Println(Number)
	fmt.Printf("Number is of type %T \n", Number)

	var Nomber int64 = 12357669900008878
	fmt.Println(Nomber)
	fmt.Printf("Number is of type %T \n", Nomber)

	//default value and some aliases
	var anotherVaribale string
	fmt.Println(anotherVaribale)
	fmt.Printf("Variable is of type %T \n", anotherVaribale)

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type %T \n", anotherVar)

}
