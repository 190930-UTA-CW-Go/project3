package new

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Set these global variable to save time in each handler
var hand string
var path = "new/templates/"

var username string

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

	username = r.FormValue("username")
	fmt.Println(username)

	present := CheckForFile(username)
	if present != true {
		CreateFile(username)
	}

	temp.Execute(w, username)
}

// CheckForFile checks for a file with the name of the username in the AWS
func CheckForFile(username string) bool {
	var doesExist bool
	fmt.Println("Checking for file named", username)
	doesExist = false // True if file does not exits
	return doesExist
}

// CreateFile creates a file in AWS
func CreateFile(username string) {
	fmt.Println("Creating file in AWS for", username)
}
