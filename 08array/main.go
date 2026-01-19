package main 

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang")

	var  fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Mango"
	//skipping index 2, it will show space for that index
	fruits[3] = "Grapes"

	fmt.Println("Fruits list is:",fruits)
	fmt.Println("Length of fruits array is:",len(fruits)) //will give the size declared , not the filled size 
}