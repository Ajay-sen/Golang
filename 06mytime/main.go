package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime:=time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006")) // this is the standard way of formatting time in Go 
	//you have to remember this date to format time as being said in the documentation 

	fmt.Println(presentTime.Format("01-02-2006 Monday")) // format with day name 

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // format with time and day name 

	fmt.Println(presentTime.Format("Monday")) // only day name 

	//if you want to create your own date and time
	createdDate:= time.Date(202,time.August, 10,23,23,0,0,time.UTC) 

	fmt.Println(createdDate)
	//again you can format this date and time too using the same Format technique
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}