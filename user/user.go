package user

import (
	"log"
	"net/http"
	"text/template"
)

// Set these global variable to save time in each handler
var hand string
var path = "user/templates/"

// Dash is the handler for the user dashboard
func Dash(w http.ResponseWriter, r *http.Request) {
	hand = path + "dash.html"
	temp, err := template.ParseFiles(path)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// MakeNew is the handler for creating a new portfolio
func MakeNew(w http.ResponseWriter, r *http.Request) {
	hand = path + "new.html"
	temp, err := template.ParseFiles(path)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// Edit is the handler for editing an exitsting portfolio
func Edit(w http.ResponseWriter, r *http.Request) {
	hand = path + "edit.html"
	temp, err := template.ParseFiles(path)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}
