package grader

import (
	"log"
	"net/http"
	"text/template"
)

// Set these global variable to save time in each handler
var hand string
var path = "grader/templates/"

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

// Rate is the handler for selecting rating a portfolio
func Rate(w http.ResponseWriter, r *http.Request) {
	hand = path + "rate.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// Submit is the handler for selecting rating a portfolio
func Submit(w http.ResponseWriter, r *http.Request) {
	hand = path + "submit.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}
