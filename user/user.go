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

// The following global variables are used to define the path to the AWS file storage location
var username string
var awsPath string
var keyPath = "~/go/src/github.com/190930-UTA-CW-Go/project3/rego.pem"
var jsonPath string
var ec2User = "ec2-user@ec2-18-188-174-65.us-east-2.compute.amazonaws.com"

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

// PortStatus store information on the user's portfolio rating. This should only be viewed by the user.
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
var rating = PortStatus{}
var jsonFile string

// Dash is the handler for the user dashboard
func Dash(w http.ResponseWriter, r *http.Request) {
	hand = path + "dash.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal("Error parsing handler files:", err)
	}
	temp.Execute(w, nil)
}

// InitDash is the handler for the user dashboard, this is the first and only dashboard
func InitDash(w http.ResponseWriter, r *http.Request) {
	hand = path + "dash.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal("Error parsing handler files:", err)
	}
	username = r.FormValue("username")
	temp.Execute(w, username)
}

// MakeNew is the handler for creating a new portfolio
func MakeNew(w http.ResponseWriter, r *http.Request) {
	hand = path + "new.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal("Error parsing handler files:", err)
	}
	temp.Execute(w, nil)
}

// Submit is the handler for creating a new portfolio
func Submit(w http.ResponseWriter, r *http.Request) {
	hand = path + "submit.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal("Error parsing handler files:", err)
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

	rating.Status = "UNCHECKED"
	rating.Comment = "UNCHECKED"

	portfolio.Information = info
	portfolio.About = about
	portfolio.Education = education
	portfolio.Project = project
	portfolio.PortStatus = rating

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
	jsonFile = username + ".json"

	hand = path + "edit.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal("Error parsing handler files:", err)
	}
	// Opens the .json file.
	readjson, readerr := os.Open(jsonFile)
	if readerr != nil {
		log.Fatal("Error opening json files:", readerr)
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
	jsonFile = username + ".json"

	// Read the json file
	portFile, err := os.Open(jsonFile)
	if err != nil {
		log.Fatal("Error opening json files:", err)
	}
	defer portFile.Close()
	byteValue, _ := ioutil.ReadAll(portFile)

	// Unmarshal the json file and save the values to the the portfolio structure
	json.Unmarshal(byteValue, &portfolio)

	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal("Error parsing handler files:", err)
	}
	temp.Execute(w, portfolio)
}

// Status is the handler for checking the status of your portfolio
func Status(w http.ResponseWriter, r *http.Request) {
	jsonFile = username + ".json"

	hand = path + "status.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal("Error parsing handler files:", err)
	}

	// Opens the .json file.
	readjson, readerr := os.Open(jsonFile)
	if readerr != nil {
		log.Fatal("Error opening json files:", readerr)
	}
	defer readjson.Close()

	// Reads the .json file.
	jsonvalue, _ := ioutil.ReadAll(readjson)

	// Grabs the info out of the .json file.
	json.Unmarshal(jsonvalue, &portfolio)
	temp.Execute(w, portfolio)
}

// Upload is the handler for uploading your portfolio
func Upload(w http.ResponseWriter, r *http.Request) {
	hand = path + "upload.html"
	awsPath = ":Portfolios/" + username
	jsonPath = "~/go/src/github.com/190930-UTA-CW-Go/project3/" + username + ".json"

	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal("Error parsing handler files:", err)
	}
	temp.Execute(w, nil)

	// The following builds a bash command and executes it
	exec.Command("bash", "-c", ("scp -i " + keyPath + " " + jsonPath + " " + ec2User + awsPath)).Run()
}
