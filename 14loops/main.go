package main 

import "fmt"

func main(){
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	//first syntax
	for  d:=0 ;d<len(days);d++{ //there is no ++d in go , its always d++
		fmt.Println(days[d])
	}


	//second syntax
	for i:=range days{
		fmt.Println(days[i])
	}


	//best one
	for index,day:= range days{
		fmt.Printf("index is %v and the value is %v\n",index,day)
	}


	roughValue := 1 

	for roughValue < 10 {

		if roughValue == 5 {
			break	//break statement
		}

		if roughValue == 3 {
			roughValue++
			continue	//continue statement
		}

		if roughValue == 2 {
			goto ajay	//goto statement
		}

		fmt.Println("Value is:",roughValue)
		roughValue++
	}

	ajay: //creating a albel to get jumped using goto statement
	    fmt.Println("Jumped from go to statement at value 2")
}