package main

import (
	"encoding/json"
	"fmt"
)


type course struct{
	Name string `json:"coursename"` //this act as a alisas while converting to json
	Price int
	Platform string `json:"websites"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"` //if the value is empty then it will omit this field
}

func main() {

	fmt.Println("welcome!")
	EncodeJson()
}

func EncodeJson(){
	
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299 , "LearnCodeOnline.in", "abc123", []string {"web -dev","js"} },
		{"NextJS Bootcamp", 199 , "LearnCodeOnline.in", "bcd123", []string {"web -dev","js"} },
		{"Angular", 299 , "LearnCodeOnline.in", "abc456", nil},
	}

	//package this data as JSON data

	finalJson, err :=json.MarshalIndent(lcoCourses,"","\t")
	if err !=nil{
		panic(err)
	}

	fmt.Printf("%s\n",finalJson)
}