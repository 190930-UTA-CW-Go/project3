package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	fileName    string
	fullURLFile string
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type inputData struct {
	Name  string
	Title string
}

func main() {
	fullURLFile = "https://raw.githubusercontent.com/jeinostroza/test/master/inputdata.json"
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
