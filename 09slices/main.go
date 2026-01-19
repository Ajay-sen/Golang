package main

import (
	"fmt"
	"sort"
)

func main() {
	//In slices we don't need to define the size of the array
	fmt.Println("Slices in Go")

	var fruitlist = []string{"Apple", "Mango", "Orange"}

	fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	//In slices we can add as many items as we want
	// it will expand the memory as needed

	//to add new item 
	fruitlist=append(fruitlist,"Mango","Banana")
	fmt.Println(fruitlist)

	//Use of colon: for range
	fruitlist1:=append(fruitlist[1:])
	fmt.Println(fruitlist1)

	fruitlist2 :=append(fruitlist[1:3])
	fmt.Println(fruitlist2)


	//another type of declaration
	highScores :=make([]int,4)
	highScores[0]=234
	highScores[1]=456
	highScores[2]=678
	highScores[3]=890
	//highScores[4]=1000 //this will give error as we have defined size 4
	//But if you use append and do the same t will reallocate the memory and add them

	highScores=append(highScores,3000,2000,1000)
	fmt.Println(highScores) 

	//sorting the values
	//there are plenty of methods that slice provide , but are not there in array

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores) 
	fmt.Println(sort.IntsAreSorted(highScores)) 


	//how to remove a value from slice based on index
	var courses = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println(courses)
	var index int=2
	courses=append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
	
}