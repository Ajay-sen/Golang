package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Welcomes to files in go")

	content := "This needs to go in a file and  be saved" 

	file, err:= os.Create("./myfile.txt") //creating the file
	if err !=nil{
		panic(err)
	}

	length, err:= io.WriteString(file, content) //writing content to the file
	if err != nil{
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close() //closing the file


	readFile("./myfile.txt") //reading the file
}



func readFile(filename string){
	databyte,err:= ioutil.ReadFile(filename)

	if err !=nil{
		panic(err)
	}

	fmt.Println("Text data inside file is \n", string(databyte))
}