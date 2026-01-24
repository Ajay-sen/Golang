package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //usually pointer

func main() {
	// greeter("Hello")
	// greeter("World")  // it will simple first print 5 times hello then 5 times world

	// go greeter("Hello")
	// greeter("world") // if you do not wait for the tread to come anf print Hello. it will just print the world 5 times and end


		websitelist := []string{
			"https://lco.dev",
			"https://go.dev",
			"https://google.com",
			"https://fb.com",
			"https://github.com",
		}
	
		for _,web := range websitelist{
			go getStatusCode(web) // by using go , we will be throwing separate thread for diff sites , which will make the things fasterv , but  then we need to use diff package for waiting else  there will be nothing printed 
			wg.Add(1)
		}
		wg.Wait() // so this will ensure , to not end the process , till the threads come back, and go will ensure separate thread call for diff sites to make the things faster

}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3* time.Second)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string){
	defer wg.Done()

	res, err :=http.Get(endpoint)

	if err != nil{
		fmt.Println("OOPS int endpoint")
	}else{
		fmt.Printf("%d status code for %s", res.StatusCode, endpoint)
	}

	
}