package main

import "fmt"

func main() {

	fmt.Println("Welcome to functions in go")
	greeter()

	result:=adder(50,78)
	fmt.Println("Result is:",result)


	proRes,message:=proadder(2,5,6,7,8)
	fmt.Println("Proadder result is:",proRes)
	fmt.Println("The message is:",message)
}

func adder(valone int,valtwo int)int{
	return valone+valtwo
}

func greeter(){
	fmt.Println("Namstey from golang")
}

func proadder(values ...int)(int,string){ // we don't know how many values are going to come , hence use treee dots syntax
	total:=0

	for _,val:=range values{
		total+=val
	}
	return total,"Hi its proadder"
}