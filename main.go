package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("To begin: open localhost:9000")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/user", userPage)
	http.HandleFunc("/grader", graderPage)
}

// homePage serves as the handler for the entry point to navigate the entire application
func homePage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// userPage serves as the handler for navigating user functionality
func userPage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/user.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

// graderPage serves as the handler for navigating grader functionality
func graderPage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/grader.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}
