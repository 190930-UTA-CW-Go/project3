package new

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
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

	//should not trigger until user enters information currently triggers immediately
	//using html it can be forced to wait however it will not leave the page
	//fixed by adding a new button and gained the ability to skip dash
	username = r.FormValue("username")
	if username != "" {
		present := CheckForFile(username)
		if present != true {
			CreateFile(username)
		} else {
			fmt.Println("User already exists please try again") //does nothing other than say this in cli
		}
	}
	temp.Execute(w, nil)
}

// CheckForFile checks for a file with the name of the username in the AWS
func CheckForFile(username string) bool {
	var doesExist bool
	fmt.Println("Checking for file named", username)
	//currently not properly connecting to database
	exec.Command("ssh -i rego.pem ec2-user@ec2-18-188-174-65.us-east-2.compute.amazonaws.com", "if [ -f ", username, " ] then doesExist=true else doesExist=false fi").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// doesExist = false // True if file does not exits
	// if doesExist == false {
	// 	CreateFile(username)
	// } else {
	// 	fmt.Println("File already exists")
	// }
	fmt.Println(doesExist) //testing
	return doesExist
}

// CreateFile creates a file in AWS
func CreateFile(username string) {
	fmt.Println("Creating file in AWS for", username)
	exec.Command("ssh -i rego.pem ec2-user@ec2-18-188-174-65.us-east-2.compute.amazonaws.com", "mkdir", username)
	fmt.Println("File Created")
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
