package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

//Portfolio stores all of the information pulled from the portfolio form. This is stored as global variable.
type Portfolio struct {
	Information Information
	About       About
	Education   Education
	Project     Project
	Status      string
}

//Information stores all of the relevant info on the person creating the portfolio.
type Information struct {
	Name  string `json:"Name"`
	Title string `json:"Title"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}

//About stores the text the user enters in the "About Me" field.
type About struct {
	Aboutme string `json:"Aboutme"`
}

//Education stores the college the user attended and the degree they got, if any.
type Education struct {
	College string `json:"College"`
	Degree  string `json:"Degree"`
}

//Project stores information on the user's projects that they have completed, if any.
type Project struct {
	Name string `json:"Name"`
	Tech string `json:"Tech"`
	Desc string `json:"Desc"`
}

// Global variables for each struct for use in different functions.
var portfolio = Portfolio{}
var info = Information{}
var about = About{}
var education = Education{}
var project = Project{}
var jsonfile string //Sets the value of the jsonfile variable.

const remoteHostname = "18.188.174.65" //IP address of amazon server
const amazon = "ec2-user@ec2-18-188-174-65.us-east-2.compute.amazonaws.com"

//This is the amazon aws hostname@IP address.
const key = "rego.pem" //This is security key used to login/SSH

/////////////////////////////////////IMPORTANT/////////////////////////////////////////////////
const localuser = "AwYiss" //USER'S LOGIN NAME. HARDCODED FOR NOW
/////////////////////////////////////IMPORTANT/////////////////////////////////////////////////

func main() {
	//http.Handle("/", http.FileServer(http.Dir("."))) Nathan's code. Dont know if its important
	http.HandleFunc("/portfolioform", portfolioform)
	http.HandleFunc("/", index)

	http.HandleFunc("/formsubmitted", formsubmitted)
	http.HandleFunc("/editportfolio", editportfolio)
	http.HandleFunc("/chooseportfolio", chooseportfolio)
	fmt.Println("Open localhost:8080")
	http.ListenAndServe(":8080", nil)
}

//index runs the index page
func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("templates/index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

//portfolioform function displays the portfolio form at localhost:8080/portfolioform
func portfolioform(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("templates/portfolioform.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}

/*formsubmitted function lets the user know their portfolio was submitted successfully.
Also automatically uploads/overwrites the user portfolio onto the amazon server*/
func formsubmitted(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("templates/formsubmitted.html")
	if err != nil {
		panic(err)
	}

	//Take inputted fullname from portfolio and use it to name the json file.
	info.Name = request.FormValue("fullname") //Part of information structure
	filename := info.Name
	filename = strings.Replace(filename, " ", "_", -1) //Changes white space to underscores
	filename = filename + ".json"                      //new filename always has underscores instead of wyt space

	//Removes file if it previously existed, to "overwrite" instead of append.
	//If the file existed before, without this code, information would be appended to old file.
	remote3 := exec.Command("rm", filename)
	remote3.Run()

	remote5 := exec.Command("mv", filename, "./portfolios")
	remote5.Run()

	//Creates and opens a new json file in the user's inputted fullname.
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}
	defer f.Close()

	//Pull information from HTML form, set to global variables.
	info.Title = request.FormValue("jobtitle")
	info.Email = request.FormValue("email")
	info.Phone = request.FormValue("phone")
	about.Aboutme = request.FormValue("aboutme")
	education.College = request.FormValue("college")
	education.Degree = request.FormValue("degree")
	project.Name = request.FormValue("projectname")
	project.Tech = request.FormValue("techused")
	project.Desc = request.FormValue("projectdesc")
	portfolio.Information = info
	portfolio.About = about
	portfolio.Education = education
	portfolio.Project = project

	portfolio.Status = "UNCHECKED" /*	This is not seen by the user in the HTML interface.
		While the user can manually edit the json file using
		Notepad, there is no other way for the user to interact
		with this field variable. Admins can change this variable
		to "ACCEPTED" or "DENIED" using the Approve-Deny-Portfolio
		program. */
	//Json file formatting
	b, err := json.MarshalIndent(portfolio, "", "    ")
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}
	f.Write(b)
	f.Close()
	http.ServeFile(response, request, "templates/formsubmitted.html")

	//These show in the console that the program has received the information entered in the form.
	fmt.Println(portfolio.Information)
	fmt.Println(portfolio.About)
	fmt.Println(portfolio.Education)
	fmt.Println(portfolio.Project)
	fmt.Println(portfolio)
	temp.Execute(response, portfolio)

	//make user portfolio on amazon. If it already exists, nothing happens
	remote4 := exec.Command("ssh", "-i", key, amazon, "mkdir", "-p", "Portfolios/"+localuser)
	remote4.Run()
	//move file to amazon server. If it already exist, gets overwritten
	remote2 := exec.Command("scp", "-i", key, filename, amazon+":Portfolios/"+localuser)
	remote2.Run()
}

/*editportfolio function takes info from an existing .json file and displays it in the form
at localhost:8080/editportfolio. This edits the LOCAL copy of the portfolio. Then upon
hitting the submit button again, BOTH local copy here and the remote copy on amazon
are overwritten with the new information.*/
func editportfolio(response http.ResponseWriter, request *http.Request) {
	//Sets the value of the jsonfile variable.
	jsonfile = "potfolios/tony.json"

	temp, err := template.ParseFiles("templates/editportfolio.html")
	if err != nil {
		panic(err)
	}
	jsonfile = request.FormValue("filename") //request filename to be modifed by the user

	//Opens the .json file.
	readjson, readerr := os.Open(jsonfile)
	if readerr != nil {
		panic(readerr)
	}
	defer readjson.Close()

	//Reads the .json file.
	jsonvalue, _ := ioutil.ReadAll(readjson)

	//Grabs the info out of the .json file.
	json.Unmarshal(jsonvalue, &portfolio)

	//Prints the values of variables in the console to make sure the .json file was read properly.
	fmt.Println(portfolio)
	fmt.Println(portfolio.About.Aboutme)
	temp.Execute(response, portfolio)
}

//chooseportfolio function prompts user to type in portfolio filename to edit.
func chooseportfolio(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("templates/chooseportfolio.html")
	if err != nil {
		panic(err)
	}

	//Requesting filename information from HTML form
	jsonfile = request.FormValue("filename") //Part of information structure of profile.
	temp.Execute(response, portfolio)
}
