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

// Model for course - file
type Course struct {
	CourseId    string  `json:"couseid"`
	CourseName  string  `json:"cousename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {

	fmt.Println("Server is running")
	r := mux.NewRouter()

	myCourses := Course{
		CourseId:    "2",
		CourseName:  "ReactJs",
		CoursePrice: 299,
		Author:      &Author{Fullname: "Pradumna", Website: "test.com"},
	}
	myCourses1 := Course{
		CourseId:    "4",
		CourseName:  "NodeJs",
		CoursePrice: 199,
		Author:      &Author{Fullname: "Jack", Website: "test.com"},
	}
	courses = append(courses, myCourses)
	courses = append(courses, myCourses1)

	//routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCouse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCouse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":4002", r))
}

// fake DB

var courses []Course

// middleware

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

// controller - file

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is working fine 🚀."))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Couse")
	w.Header().Set("Content-Type", "application/json")

	// get ID from req
	params := mux.Vars(r)

	// loop though the couses, find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with the given id")

}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Couse")
	w.Header().Set("Content-Type", "application/json")

	// If bosy is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Plesae send some data")

	}

	// If the bosy is {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//genrate unique id string

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCouse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one Couse")
	w.Header().Set("Content-Type", "application/json")

	// first - grab ID fom req
	params := mux.Vars(r)

	// loop, id, remove, add with the ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // remove the element at index 2. :index is from 0 to index-1, index+1: is from index+1 to end

			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"] // Overidding the value
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Id is not found")
}

func deleteOneCouse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one Couse")
	w.Header().Set("Content-Type", "application/json")

	// first - grab ID fom req
	params := mux.Vars(r)

	// loop, id, remove,

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // remove the element at index 2. :index is from 0 to index-1, index+1: is from index+1 to end
			json.NewEncoder(w).Encode("Data with ID deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("Id is not found")
}
