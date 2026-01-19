package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input , _:=reader.ReadString('\n')

	fmt.Println("Thanks for rating us ", input)

	//we know that input is string but we want to convert it to integer 

	numRating,err :=strconv.ParseFloat(strings.TrimSpace(input),64)
	
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your Rating: ", numRating+1)
	}
}