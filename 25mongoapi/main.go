package main

import (
	"fmt"
	"net/http"
	"github.com/ajay-sen/mongoapi/router"
)

func main() {
	fmt.Println("MondoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening at port 4000 ...")
} 