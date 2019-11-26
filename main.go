package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/190930-UTA-CW-Go/project3/grader"
	"github.com/190930-UTA-CW-Go/project3/new"
	"github.com/190930-UTA-CW-Go/project3/user"
)

// Set these global variable to save time in each handler
var hand string
var path = "templates/"

func main() {
	// The following handlers are located inside the main package.
	http.HandleFunc("/", homePage)
	http.HandleFunc("/user", userPage)
	http.HandleFunc("/grader", graderPage)

	// The following handlers are located inside the user package.
	http.HandleFunc("/user/dash", user.Dash)
	http.HandleFunc("/user/new", user.MakeNew)
	http.HandleFunc("/user/edit", user.Edit)

	// The following handlers are located inside the grader package.
	http.HandleFunc("/grader/dash", grader.Dash)

	// The following handlers are located inside the new package.
	http.HandleFunc("/new", new.User)
	http.HandleFunc("/new/dash", new.Dash)

	// Open up and start listening!
	fmt.Println("To begin: open localhost:9000")
	http.ListenAndServe(":9000", nil)
}

// homePage serves as the handler for the entry point to navigate the entire application
func homePage(w http.ResponseWriter, r *http.Request) {
	hand = path + "index.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// userPage serves as the handler for navigating user functionality
func userPage(w http.ResponseWriter, r *http.Request) {
	hand = path + "user.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// graderPage serves as the handler for navigating grader functionality
func graderPage(w http.ResponseWriter, r *http.Request) {
	hand = path + "grader.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// The following global variables are used to format protfolios

// Portfolio stores all of the information pulled from the portfolio form. This is stored as global variable.
type Portfolio struct {
	Information Information
	About       About
	Education   Education
	Project     Project
	Status      string
}

// Information stores all of the relevant info on the person creating the portfolio.
type Information struct {
	Name  string `json:"Name"`
	Title string `json:"Title"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}

// About stores the text the user enters in the "About Me" field.
type About struct {
	Aboutme string `json:"Aboutme"`
}

// Education stores the college the user attended and the degree they got, if any.
type Education struct {
	College string `json:"College"`
	Degree  string `json:"Degree"`
}

// Project stores information on the user's projects that they have completed, if any.
type Project struct {
	Name string `json:"Name"`
	Tech string `json:"Tech"`
	Desc string `json:"Desc"`
}

// These global variables for each struct for use in different functions.
var portfolio = Portfolio{}
var info = Information{}
var about = About{}
var education = Education{}
var project = Project{}
var jsonfile string
