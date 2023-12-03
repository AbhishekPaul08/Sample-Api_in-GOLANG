package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Courses struct {
	CourseId    string   `json:"courseid"`
	CourseName  string   `json:"cousename"`
	CoursePrice int      `json:"price"`
	Author      *Authors `json:"author"`
}

type Authors struct {
	AuthorName    string `json:"name"`
	AuthorWebsite string `json:"website"`
}

var course []Courses //for fake  db

func (c *Courses) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
func main() {
	fmt.Println("Success")
	r := mux.NewRouter()
	course = append(course, Courses{CourseId: "1", CourseName: "OS", CoursePrice: 233, Author: &Authors{AuthorName: "karthik", AuthorWebsite: "Aba.in"}})
	course = append(course, Courses{CourseId: "2", CourseName: "DBMS", CoursePrice: 1000, Author: &Authors{AuthorName: "Amith", AuthorWebsite: "Amiths.in"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllcouses).Methods("GET")
	r.HandleFunc("/course/{id}", getOnecouse).Methods("GET")
	r.HandleFunc("/course", createOnecourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOne).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// controllers below
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to new Server</h1>"))
}

func getAllcouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func getOnecouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) //to grab id

	// checking if id is valid
	for _, c := range course {
		if c.CourseId == params["id"] {
			json.NewEncoder(w).Encode(c)
			return
		}

	}
	json.NewEncoder(w).Encode(params["id"] + " not found")

}

func createOnecourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil { //to check if body is nil
		json.NewEncoder(w).Encode("Send valid data")
	}

	//if data is being sent like {}
	var cour Courses
	_ = json.NewDecoder(r.Body).Decode(&cour)

	// if cour.IsEmpty() {
	// 	json.NewEncoder(w).Encode("Empty")
	// 	return
	// }
	// for _, courr := range course {
	// 	if courr.CourseName == cour.CourseName {
	// 		json.NewEncoder(w).Encode("REpaet")
	// 		break
	// 		return

	// 	}
	// 	cour.CourseId = strconv.Itoa(rand.Intn(100))
	// 	course = append(course, cour)
	// 	json.NewEncoder(w).Encode(cour)
	// 	return
	// }
	// Check if the course already exists
	for _, courr := range course {
		if courr.CourseName == cour.CourseName {
			json.NewEncoder(w).Encode("Repeat")
			return
		}
	}

	// Generate a new CourseId and add the course
	cour.CourseId = strconv.Itoa(rand.Intn(100))
	course = append(course, cour)
	json.NewEncoder(w).Encode(cour)

}

func updateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//loop grab delete and update

	for index, couse := range course {
		if couse.CourseId == params["id"] {
			course = append(course[:index], course[index+1:]...)
			var couse Courses
			_ = json.NewDecoder(r.Body).Decode(&couse)
			couse.CourseId = params["id"]
			course = append(course, couse)
			json.NewEncoder(w).Encode(couse)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	//loop, grab delete
	for index, cour := range course {
		if cour.CourseId == params["id"] {
			course = append(course[:index], course[index+1:]...)
			json.NewEncoder(w).Encode("Deleted successfully.")
			break
		} else {
			json.NewEncoder(w).Encode(params["id"] + " not found")
			break
		}
	}
}
