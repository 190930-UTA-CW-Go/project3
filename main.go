package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//Portfolio stores all of the information pulled from the portfolio form. This is stored as global variable.
type Portfolio struct {
	Information Information
	About       About
	Education   Education
	Project     Project
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

/*Global variables for each struct for use in different functions.*/

//Stores the Portfolio struct, which stores all other structs.
var portfolio = Portfolio{}

//Stores the Infomation struct.
var info = Information{}

//Stores the About struct.
var about = About{}

//Stores the Education struct.
var education = Education{}

//Stores the Project struct.
var project = Project{}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/portfolioform", portfolioform)
	http.HandleFunc("/formsubmitted", formsubmitted)
	fmt.Println("Open localhost:8080")
	http.ListenAndServe(":8080", nil)
}

//The portfolioform function displays the portfolio form at localhost:8080/portfolioform
func portfolioform(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("html/portfolioform.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}

//The formsubmitted function lets the user know their portfolio was submitted successfully.
func formsubmitted(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("html/formsubmitted.html")
	if err != nil {
		panic(err)
	}

	//Take inputted fullname from profile and use it to name the json file
	info.Name = request.FormValue("fullname") //Part of information structure of profile.
	filename := info.Name + ".json"

	//Creates and opens a new json file in the user's inputted fullname.
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}
	defer f.Close()

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

	b, err := json.MarshalIndent(portfolio, "", "    ")
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}
	f.Write(b)
	f.Close()
	http.ServeFile(response, request, "html/formsubmitted.html")

	//These show in the console that the program has received the information entered in the form.
	fmt.Println(portfolio.Information)
	fmt.Println(portfolio.About)
	fmt.Println(portfolio.Education)
	fmt.Println(portfolio.Project)
	fmt.Println(portfolio)

	temp.Execute(response, portfolio)

}
