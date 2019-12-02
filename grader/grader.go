package grader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/190930-UTA-CW-Go/project3/grader/downloadfile"
)

// Set these global variable to save time in each handler
var hand string
var path = "grader/templates/"

var username string

// Dash is the handler for the grader dashboard
func Dash(w http.ResponseWriter, r *http.Request) {
	hand = path + "dash.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// Selector is the handler for selecting a portfolio to grade
func Selector(w http.ResponseWriter, r *http.Request) {
	hand = path + "select.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// View is the handler for loading the selected portfolio
func View(w http.ResponseWriter, r *http.Request) {
	hand = path + "view.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}

	username = r.FormValue("username")
	downloadfile.DownloadAWS(username)

	temp.Execute(w, username)
}

// Rate is the handler for viewing and rating the portfolio
func Rate(w http.ResponseWriter, r *http.Request) {
	// Sets the value of the jsonfile variable.
	jsonFile = username + ".json"

	hand = path + "rate.html"
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

// Submit is the handler for submitting the portfolio with the new rating
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

	remote5 := exec.Command("mv", filename, "./portfolios")
	remote5.Run()

	//Creates and opens a new json file in the user's inputted fullname.
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()

	// The following assigns the form values to the structs going into the json file
	info.Title = r.FormValue("jobtitle")
	info.Email = r.FormValue("email")
	info.Phone = r.FormValue("phone")

	about.Aboutme = r.FormValue("aboutme")

	education.College = r.FormValue("college")
	education.Degree = r.FormValue("degree")

	project.Name = r.FormValue("projectname")
	project.Tech = r.FormValue("techused")
	project.Desc = r.FormValue("projectdesc")

	portStatus.Status = r.FormValue("rating")
	portStatus.Comment = r.FormValue("comments")

	portfolio.Information = info
	portfolio.About = about
	portfolio.Education = education
	portfolio.Project = project
	portfolio.PortStatus = portStatus

	// The following packages the structs into the json file.
	b, err := json.MarshalIndent(portfolio, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	f.Write(b)
	f.Close()

	// Serve and execute
	http.ServeFile(w, r, hand)
	temp.Execute(w, portfolio)
}

// The following global variables are used to format protfolios

// Portfolio stores all of the information pulled from the portfolio form. This is stored as global variable.
type Portfolio struct {
	Information Information
	About       About
	Education   Education
	Project     Project
	PortStatus  PortStatus
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

// PortStatus stores information about user's portfolio rating.
type PortStatus struct {
	Status  string `json:"Status"`
	Comment string `json:"Comment"`
}

// These global variables for each struct for use in different packages.

// Portfolio holds the Portfolio struct.
var portfolio = Portfolio{}
var info = Information{}
var about = About{}
var education = Education{}
var project = Project{}
var portStatus = PortStatus{}
var jsonFile string
