package main

import "fmt"

func main() {

	fmt.Println("Maps in Golang!")

	//For creating a map , for key-value pairs
	languages :=make(map[string]string)

	languages["JS"]="JavaScript"
	languages["PY"]="Python"
	languages["RB"]="Ruby" 

	fmt.Println("List of all languages: ",languages)
	fmt.Println("JS shorts for: ",languages["JS"])


	//for deleting elements from map
	delete(languages,"RB")
	fmt.Println("List of languages:",languages) 


	//loops are interesting in golang
	for key, value:=range languages{
		fmt.Printf("For key %v, value is %v\n",key,value)
	}
}
