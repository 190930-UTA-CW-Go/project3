package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
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
	http.HandleFunc("/download", downloadFunc)
	http.HandleFunc("/page", portfoliopage)
	fmt.Println(" Open browser to localhost:7004")
	http.ListenAndServe(":7004", nil)

}

func downloadFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running downloadfunc")
	GettingFile("Portfolios", "Example Portfolio.json")
	temp, err := template.ParseFiles("download.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)

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

var (
	fileName     string
	fullURLFile  string
	urlPath      string
	folder       string
	fileDownload string
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type inputData struct {
	Name  string
	Title string
}

// GettingFile doc
func GettingFile(folder string, fileDownload string) {
	urlPath = "190930-UTA-CW-Go/project3/dev"
	fullURLFile = "https://raw.githubusercontent.com/" + urlPath + "/" + folder + "/" + fileDownload
	buildFileName()      // Build fileName from fullPath
	file := createFile() // Create blank file

	// inputData1 := new(inputData)
	// getJSON(fullURLFile, &inputData1)
	// println(inputData1.Name)
	// println(inputData1.Title)

	// Put content on file
	success, size := putFile(file)
	if success == true {
		fmt.Println("File downloaded as: " + fileName)
		fmt.Println("Size: " + strconv.FormatInt(size, 10))
	} else {
		fmt.Println("File not downloaded.")
	}

}

// buildFileName takes unnecessary information from the URL and builds the file name
func buildFileName() {
	fileURL, err := url.Parse(fullURLFile)
	checkError(err)

	path := fileURL.Path
	segments := strings.Split(path, "/")

	fileName = segments[len(segments)-1]
}

// checkError panics and prints error codes if they exist
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// createFile creates a blank file in local storage
func createFile() *os.File {
	file, err := os.Create(fileName)

	checkError(err)
	return file
}

// putFile puts the body of the file passed into the blank file created earlier
func putFile(file *os.File) (bool, int64) {
	client := httpClient()

	resp, err := client.Get(fullURLFile)
	checkError(err)
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	defer file.Close()
	checkError(err)

	success := true
	return success, size
}

// httpClient. Not sure what this does
func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}
