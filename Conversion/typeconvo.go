package main

import (
	"fmt"
	"reflect"
)

func main() {

	//int to float64
	var id int = 13
	fmt.Println("id is: ", id)
	fmt.Println("Type of id is:", reflect.TypeOf(id))

	new_id := float64(id)
	fmt.Println("New id is: ", new_id)
	fmt.Println("Type of new_id is:", reflect.TypeOf(new_id))

	//float to int
	var id1 float64 = 13.45
	fmt.Println("id is: ", id1)
	fmt.Println("Type of id1 is: ", reflect.TypeOf(id1))

	new_id1 := int(id1)
	fmt.Println("New id is: ", new_id1)
	fmt.Println("Type of new id is: ", reflect.TypeOf(new_id1))

	//string to []byte

	var name string = "Gitesh"
	fmt.Println("Name is: :", name)
	fmt.Println("Name type is: ", reflect.TypeOf(name))

	new_name := []byte(name)
	fmt.Println("New name is: ", new_name)
	fmt.Println("New name type is: ", reflect.TypeOf(new_name))

	//[]byte to string

	var byteData []byte = []byte{20, 21, 22, 22}
	fmt.Println("byteData is: ", byteData)
	fmt.Println("Type of byteData is: ", reflect.TypeOf(byteData))

	stringData := string(byteData)
	fmt.Println("New stringData is: ", stringData)
	fmt.Println("Type of new stringData is: ", reflect.TypeOf(stringData))

}
