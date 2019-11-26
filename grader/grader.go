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
