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
