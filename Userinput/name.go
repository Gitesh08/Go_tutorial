package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wlc := "Welcome to my world of Golang"
	fmt.Println(wlc)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your favorite number: ")

	//comma ok // error ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for input: ", input)

}
