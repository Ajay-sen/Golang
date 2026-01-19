package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the User Input Program!"
	fmt.Println(welcome)

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your rating: ")
	
	//comma ok idiom || err ok idiom 

	input, _:=reader.ReadString('\n')
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("Type of this rating is %T", input)

	//where there are chances of error use comma ok idiom that is : input,err   
	//Else we can use underscore to ignore the error value

	// and the inputs by default are of string type
}