package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghs67h4"

func main() {

	fmt.Println("welcome to urls in golang")
	fmt.Println(myurl)


	//parsing urls
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//if you want to get all the query params as a map
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n",qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])


	for _,val:=range qparams{
		fmt.Println("Param is: ",val)
	}

	//you have to build urls as well, from a number of chunks or smaller parts as below
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tuts/learn",
		RawQuery: "user=abcd",
	}

	anotherUrl:= partsOfUrl.String()
	fmt.Println(anotherUrl)
}