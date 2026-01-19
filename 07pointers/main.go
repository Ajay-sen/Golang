package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var one int=42 // normal variable
	var two *int // pointer variable  

	fmt.Println("value of pointer is:", two) // by deafult pointer is nil

	two=&one // assign address of one to pointer two

	fmt.Println("value of pointer is:", two)

	//actually with * mark it will dereference the pointer(ie will show the vale=42)
	//But with just & and no * it will show the address of the variable

	myNumber :=23
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is:",ptr)
	fmt.Println("Value of dereferenced pointer is:",*ptr) 


	//Modifying actual values via pointers
	*ptr = *ptr + 2
	fmt.Println("New Value of myNumber is:",myNumber)

}