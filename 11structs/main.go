package main 

import "fmt"

func main(){

	fmt.Println("Structs in golang")
	//No inheritence in golang, No super or base classes 

	ajay :=User{"Ajay", "ajay@gmail.com", true, 20}
	fmt.Println(ajay)
	fmt.Printf("hitest details are: %+v\n",ajay)
	fmt.Printf("Name is %v and email is %v\n",ajay.Name, ajay.Email) 

}


type User struct{
	Name string
	Email string
	Status bool 
	Age int 
}