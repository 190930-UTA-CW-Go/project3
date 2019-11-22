package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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
	Education string   s`json:"education"`
	Projects  []string `json:"projects"`*/
}

//HTML template passing in values from inputdata.json

func main() {
	http.HandleFunc("/", portfoliopage)
	fmt.Println(" Open browser to localhost:7004")
	http.ListenAndServe(":7004", nil)

}

//page creates template for portfolio
func portfoliopage(response http.ResponseWriter, request *http.Request) {
	//Opens inputdata.json file
	jsonFile, err := os.Open("inputdata.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	var values portfolioInfo

	//reads file and stores content as []byte
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &values)

	fmt.Println(values.Name)
	fmt.Println(values.Title)

	// portfolioinfo := portfolioInfo{
	// 	Name:  values.Name,
	// 	Title: values.Title,
	// }

	temp, _ := template.ParseFiles("page.html")

	temp.Execute(response, values)

}
