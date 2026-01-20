package main 


import "fmt"

func main(){

	fmt.Println("If else statement in go")
	
	loginCount:=23
	var result string

	if loginCount<10{ // in  go it is mandatory to keep these curly braces in the same line with if
		result="Regular user"
	}else if  loginCount >10{
		result = "watch out"
	}else {
		result="exactly 10 login count"
	}

	fmt.Println(result)


	//cheching on the going number is even or odd
	if 9%2==0{
		fmt.Println("even number")
	}else{
		fmt.Println("odd number")
	}


	//special syntax to be used in web handling

	if num:=3; num<10{
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Number is greater than equalto  10")
	}


	//for handling err is nil or not:
	// if err !=nill{}

}