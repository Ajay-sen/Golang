package main

import "fmt"

const LoginToekn string="qwertyuiop" // public variable, since starting with capital letter

func main() {
	// fmt.Println("Hello, Variables!")
	var username string="Ajay"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool=true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	//you may read about a lots of other data types in official docs
	// https://golang.org/ref/spec#Types


	//Implicit type variable declaration
	// lexer infers the type by the value assigned
	var userAge=25
	fmt.Println(userAge)
	fmt.Printf("Variable is of type: %T \n",userAge)


	// no var style declaration
	//you can use := operator just inside a function
	country:="India"
	fmt.Println(country)
	fmt.Printf("Variable is of type: %T \n",country)
}