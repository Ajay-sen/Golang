package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for course - <separate file>
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	website string `json:"website"`
}

//fake DB
var courses []Course

//middleware, helper --> separate file
func (c *Course) IsEmpty() bool{
	// return c.CourseId =="" && c.CourseName ==""
	return c.CourseName ==""
}


func main(){
	fmt.Println("API- learnCodeonline.in")

	r :=mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2",CourseName:"ReactJS", CoursePrice: 299, Author: &Author{Fullname:"Ajay Sen", website:"lco.dev"}})

	courses = append(courses, Course{CourseId: "4",CourseName:"NextJS", CoursePrice: 599, Author: &Author{Fullname:"Rahul gupta", website:"go.dev"}})

	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")


	//listen to a port
	log.Fatal(http.ListenAndServe(":4000",r))
}


//controllers -> separate file


//serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API by learnCourseOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r*http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}


func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type","application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses , find matching id and return the response
	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id") // if no course found with the given id
	return
}


func createOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("create one course")
	w.Header().Set("Content-Type","application/json")

	//what if:body is empty
	if r.Body ==nil{
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}
	var course Course
	_=json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data insidie JSON")
		return
	}

	//generate unique ID, convert that to string
	//append course into the Courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses=append(courses,course)
	json.NewEncoder(w).Encode(course)
	return 
}


func updateOneCourse (w http.ResponseWriter, r *http.Request){

	fmt.Println("Create one course")
	w.Header().Set("Content-Type","application/json")

	//first - grab id from req
	params := mux.Vars(r)
	//loop through the values, find the id that need to be updated, remove  that from the array/slice, add with my ID newdata updated

	for index, course := range courses {
		if course.CourseId ==params["id"]{
			courses =append(courses[:index],courses[index+1:]...)

			_=json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			courses=append(courses,course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//if id is not found 
	json.NewEncoder(w).Encode("Id does not exist")
	return
}


func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type","application/json")

	params :=mux.Vars(r)

	//loop the datas , found the id ,remove (index, index+1)

	for index, course:=range courses {
		if course.CourseId ==params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			break
		}
	}

	//if id does not match 
	json.NewEncoder(w).Encode("Given Id does Exist for deleting")
}