package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web request verbs")
	//PerformGetRequest() //remember to get the output you need to start the node js server , kept in the webserver folder

	//PerformPostJsonRequest()

	PerformPostFormRequest()
}



//Get request handling function
func PerformGetRequest(){

	const myurl = "http://localhost:8000/get" //instead of this consttant lonk , you can take input this link as a string from user

	response , err:=http.Get(myurl)
	if err !=nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("STatus code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	//METHOD 1:
	// fmt.Println("Content is: ", content) will give the output in Byte format

	fmt.Println(string(content)) // converting byte format to string format

	//METHOD 2:
	//ANOTHER WAY OF DOING THIS IN GO LANGUAGE using strings buider method
	var responseString strings.Builder 
	byteCount, _:=responseString.Write(content)

	fmt.Println("ByteCount is:",byteCount)
	fmt.Println(responseString.String())

	//you can use any of the methods to get the output
}



// POST request handling function
func PerformPostJsonRequest(){

	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Lets go with golang",
			"price":0,
			"platform":"learncodeonline.in"
		}
	`)
	response, err:=http.Post(myurl, "application/json",requestBody)

	if err !=nil{
		panic(err)
	}

	defer response.Body.Close()

	content ,_ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}



//handling post form requests

func PerformPostFormRequest(){

	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname","ajay")
	data.Add("lastname","sen")
	data.Add("email","ajay@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}