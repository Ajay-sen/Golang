package main 

import "fmt"

func main(){

	fmt.Println("Structs in golang")
	//No inheritence in golang, No super or base classes 

	ajay :=User{"Ajay", "ajay@gmail.com", true, 20}
	fmt.Println(ajay)
	fmt.Printf("hitest details are: %+v\n",ajay)
	fmt.Printf("Name is %v and email is %v\n",ajay.Name, ajay.Email) 

	ajay.GetStatus()
	ajay.NewMail() //but it wil not change the email as we are not using pointer receiver
	fmt.Println(ajay.Email)

}


type User struct{
	Name string
	Email string
	Status bool 
	Age int 
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "testgo.dev"
	fmt.Println("Email of this user is : ", u.Email)
}