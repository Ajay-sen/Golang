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
	// EncodeJson()
	DecodeJson()
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



// decoding json
func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS",
		"Price": 299,
		"websites": "learncodeonline.in",
		"tags": ["web-dev","js"]
	}
	
	`)

	var lcoCourse course

	checkValid :=json.Valid(jsonDataFromWeb) //check if the json data came is valid or not
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	}else{
		fmt.Println("JSON WAS NOT VALID")
	}


	// some cases where you just want to  add data to key value pairs
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for k,v:= range myOnlineData{
		fmt.Printf("Key is %v and value is %v and Type is: %T\n",k,v,v)
	}
}