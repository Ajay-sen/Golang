package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two") // defer works in LIFO order, so after main function ends, it will print Two, One, World

	fmt.Println("Hello")
	myDefer()

}


func myDefer(){

	for i:=0;i<5;i++{
		defer fmt.Print(i)
	}
}