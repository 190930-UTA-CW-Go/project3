package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/portfolioform", portfolioform)
	http.HandleFunc("/formsubmitted", formsubmitted)
	fmt.Println("Open localhost:8080")
	http.ListenAndServe(":8080", nil)
}

//The portfolioform function displays the portfolio form at localhost:8080/portfolioform
func portfolioform(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("templates/portfolioform.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}

//The formsubmitted function lets the user know their portfolio was submitted successfully.
func formsubmitted(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("templates/formsubmitted.html")
	if err != nil {
		panic(err)
	}

	//Take inputted fullname from profile and use it to name the json file.
	info.Name = request.FormValue("fullname") //Part of information structure of profile.
	filename := info.Name
	filename = strings.Replace(filename, " ", "_", -1) //Changes white space to underscores
	filename = filename + ".json"

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
	portfolio.Status = "UNCHECKED"

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
}
