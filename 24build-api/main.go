package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// creating a model for courses - file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrize int     `json:"prize"`
	Author      *Author `json:"author"`
	//pointer to struct Author since we do not want to pass all the vars here
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file

// returns true if both courseName and CourseId both are empty
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	// reader(r) is from where we get the response
	// writer(w) is wehre we write the response, i.e., headers, etc. stuff
	w.Write([]byte("<h1>Welcome to API by learnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	// set the header of the response
	json.NewEncoder(w).Encode(courses)
	// whatever we pass inside encode function will be treated as a json value &
	// will be thrown back to whoever is making a request via this `w`
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	// max will gather all the var params from the req url, we can do that with url as well
	fmt.Printf("params : \n%+v\n", params)

	// loop through courses, find matching id & return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("empty data sent: body is nil")
	}

	// id body = {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("send some data: empty json data")
	}

	// at this point the data should be correct
	// rand.Seed(time.Now().Unix())
	// rand.New(rand.NewSource())
}
