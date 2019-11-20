package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// portfolioInfo Structs should represent the JSON string
type portfolioInfo struct {
	Name  string `json:"Name"`
	Title string `json:"Title"`

	//These are extra struct field variables I excluded to keep this program simple.
	/*Email     string   `json:"email"`
	Phone     int      `json:"phone"`
	Aboutme   []string `json:"aboutme"`
	Education string   `json:"education"`
	Projects  []string `json:"projects"`*/
}

//open is the "main page" of localhost:8080
func open(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("portfolio.html")
	t.Execute(w, nil)
}

//portfolio handles the user input from the html page. Converts it to a json file named "inputdata.json"
func portfolio(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile("inputdata.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()

	userInput := new(portfolioInfo)
	userInput.Name = r.FormValue("name")
	userInput.Title = r.FormValue("title")

	b, err := json.Marshal(userInput)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	f.Write(b)
	f.Close()
	http.ServeFile(w, r, "portfolio.html")
}

func main() {
	http.HandleFunc("/portfolio", portfolio)
	http.HandleFunc("/", open)
	fmt.Println("Open localhost:8080")
	http.ListenAndServe(":8080", nil)
}
