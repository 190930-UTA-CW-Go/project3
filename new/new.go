package new

import (
	"html/template"
	"log"
	"net/http"
)

// Set these global variable to save time in each handler
var hand string
var path = "new/templates/"

// User is the handler for creating new users
func User(w http.ResponseWriter, r *http.Request) {
	hand = path + "user.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// Dash is the handler for navigating new users back to the user login
func Dash(w http.ResponseWriter, r *http.Request) {
	hand = path + "dash.html"
	temp, err := template.ParseFiles(hand)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}
