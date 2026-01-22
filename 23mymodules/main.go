package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000",r))
}

func greeter(){
	fmt.Println("hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to golang</h1>"))
} 


// Few important commands:
// go mod init <module-name>  --> to initialize a module
// go mod tidy  --> to add missing and remove unused modules
// go get <module-name>  --> to install a new module
// go mod download  --> to download all the modules to local system
// go mod verify  --> to verify the dependencies have expected content
// go list -m all  --> to list all the modules in the project
//go list -m -version github.com/gorilla/mux  --> to check the version of a module
//go mod why github.com/gorilla/mux  --> to know why a module is needed
// go mod graph  --> to see the module dependency graph
// go mod vendor  --> to create a vendor directory with all dependencies
// go run -mod=vendor main.go  --> to run the code using vendor directory(if there the module is available in vendor directory), else will bring from web