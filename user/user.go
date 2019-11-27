package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)

// Set these global variable to save time in each handler
var hand string
var path = "user/templates/"

var username string

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

// These global variables for each struct for use in different packages.

// Portfolio holds the Portfolio struct.
var portfolio = Portfolio{}
var info = Information{}
var about = About{}
var education = Education{}
var project = Project{}
var jsonFile string

// Dash is the handler for the user dashboard
func Dash(w http.ResponseWriter, r *http.Request) {
	hand = path + "dash.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	username = r.FormValue("username")
	temp.Execute(w, username)
}

// MakeNew is the handler for creating a new portfolio
func MakeNew(w http.ResponseWriter, r *http.Request) {
	hand = path + "new.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// Submit is the handler for creating a new portfolio
func Submit(w http.ResponseWriter, r *http.Request) {
	hand = path + "submit.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}

	//Take inputted fullname from profile and use it to name the json file.
	info.Name = r.FormValue("fullname") //Part of information structure of profile.
	filename := username + ".json"

	//Removes file if it previously existed, to "overwrite" instead of append.
	//If the file existed before, without this code, information would be appended to old file.
	remote3 := exec.Command("rm", filename)
	remote3.Run()

	remote5 := exec.Command("mv", username, "./portfolios")
	remote5.Run()

	//Creates and opens a new json file in the user's inputted fullname.
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()

	info.Title = r.FormValue("jobtitle")
	info.Email = r.FormValue("email")
	info.Phone = r.FormValue("phone")

	about.Aboutme = r.FormValue("aboutme")

	education.College = r.FormValue("college")
	education.Degree = r.FormValue("degree")

	project.Name = r.FormValue("projectname")
	project.Tech = r.FormValue("techused")
	project.Desc = r.FormValue("projectdesc")

	portfolio.Information = info
	portfolio.About = about
	portfolio.Education = education
	portfolio.Project = project
	portfolio.Status = "UNCHECKED"

	b, err := json.MarshalIndent(portfolio, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	f.Write(b)
	f.Close()
	http.ServeFile(w, r, hand)
	temp.Execute(w, portfolio)
}

// Edit is the handler for editing an exitsting portfolio
func Edit(w http.ResponseWriter, r *http.Request) {
	// Sets the value of the jsonfile variable.
	jsonFile = "Tony_Moon.json"

	hand = path + "edit.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	// Opens the .json file.
	readjson, readerr := os.Open(jsonFile)
	if readerr != nil {
		panic(readerr)
	}
	defer readjson.Close()

	// Reads the .json file.
	jsonvalue, _ := ioutil.ReadAll(readjson)

	// Grabs the info out of the .json file.
	json.Unmarshal(jsonvalue, &portfolio)
	temp.Execute(w, portfolio)
}

// Printer is the handler for printing your portfolio
func Printer(w http.ResponseWriter, r *http.Request) {
	hand = path + "print.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// Status is the handler for checking the status of your portfolio
func Status(w http.ResponseWriter, r *http.Request) {
	jsonFile = username + ".json"

	hand = path + "status.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)

	// Opens the .json file.
	readjson, readerr := os.Open(jsonFile)
	if readerr != nil {
		panic(readerr)
	}
	defer readjson.Close()
}
