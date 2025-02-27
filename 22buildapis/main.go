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

// model for coure - file (goes in file)
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// dummy db
var Courses []Course

// middleware or helper
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	Courses = append(Courses, Course{CourseId: "2",
		CourseName: "React", CoursePrice: 299, Author: &Author{FullName: "John", Website: "Learn.in"}})

	Courses = append(Courses, Course{CourseId: "4",
		CourseName: "NodeJS", CoursePrice: 399, Author: &Author{FullName: "John", Website: "Learn.in"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers -file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//get id from request
	params := mux.Vars(r)

	//return map of strings
	// fmt.Println(params)
	// fmt.Printf("params type: %T\n", params)

	//loop through courses,find matching id and return the response
	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "No course found with this id",
		"id":      params["id"],
	})
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about if data is {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the json")
		return
	}

	//check if coursename already exists
	for _, val := range Courses {
		//fmt.Println(val.CourseName, course.CourseName)
		if val.CourseName == course.CourseName {
			json.NewEncoder(w).Encode(map[string]string{
				"message": "This course already exists",
				"id":      val.CourseId,
			})
			return
		}
	}

	//generate unique id, string
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from req
	params := mux.Vars(r)

	//loop->id->remove->add with param id
	for ind, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:ind], Courses[ind+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			Courses = append(Courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//send res when id is not found
	json.NewEncoder(w).Encode(map[string]string{
		"message": "No course found with this id",
		"id":      params["id"],
	})
}

// delete one course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop->id->remove
	for ind, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:ind], Courses[ind+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Course deleted successfully",
				"id":      params["id"],
			})
			return
		}
	}

	//send json response
	json.NewEncoder(w).Encode(map[string]string{
		"message": "No course found with this id",
		"id":      params["id"],
	})
}
