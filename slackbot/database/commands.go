package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	_ "github.com/lib/pq" // necessary
)

const amazon = "ec2-user@ec2-18-188-174-65.us-east-2.compute.amazonaws.com" //Amazon aws hostname@IPaddress.
const key = "../rego.pem"                                                   //This is security key used to login/SSH

// Portfolio stores all of the information pulled from the portfolio form. This is stored as global variable.
type Portfolio struct {
	Information Information `json:"Information"`
	About       About       `json:"About"`
	Education   Education   `json:"Education"`
	Project     Project     `json:"Project"`
	PortStatus  PortStatus  `json:"PortStatus"`
}

// Information stores all of the relevant info on the person creating the portfolio.
type Information struct {
	Name  string `json:"Name"`
	Title string `json:"Title"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}

// About stores the text the user enters in the "About Me" field.
type About struct {
	Aboutme string `json:"Aboutme"`
}

// Education stores the college the user attended and the degree they got, if any.
type Education struct {
	College string `json:"College"`
	Degree  string `json:"Degree"`
}

// Project stores information on the user's projects that they have completed, if any.
type Project struct {
	Name string `json:"Name"`
	Tech string `json:"Tech"`
	Desc string `json:"Desc"`
}

// PortStatus store information on the user's portfolio rating. This should only be viewed by the user.
type PortStatus struct {
	Status  string `json:"Status"`
	Comment string `json:"Comment"`
}

// NewFolder = create folder and json file named after "username"
// If folder already exists return false else return true
func NewFolder(username string) bool {
	fmt.Println("Creating file in AWS for", username)
	err := exec.Command("bash", "-c", ("ssh -i " + key + " " + amazon + " mkdir Portfolios/" + username)).Run()
	if err != nil {
		fmt.Println("Folder already exists!")
		return false
	}
	fmt.Println("File Created")
	exec.Command("bash", "-c", ("ssh -i " + key + " " + amazon + " touch Portfolios/" + username + "/" + username + ".json")).Run()
	return true
}

// GetStatus = get date and time of last time folder was modified
// Format date and time using Go time constants (look at documentation online)
func GetStatus(username string) (bool, string) {
	output, err := exec.Command("bash", "-c", ("ssh -i " + key + " " + amazon + " stat Portfolios/" + username + " | egrep 'Modify'")).Output()
	if err != nil {
		return false, "ERROR"
	}
	fmt.Println(string(output))
	slice := strings.Fields(string(output))

	dateTime := slice[1] + " " + (slice[2])[:5]
	fmt.Println("dateTime:", dateTime)

	t, _ := time.Parse("2006-01-02 15:04", dateTime)
	result := t.Format("January 2, 2006 at 3:04 PM")
	fmt.Println(result)
	return true, result
}

// FindFile =
func FindFile(email string) (string, string) {
	hold := ParseEmail(email)

	path := "../Portfolios/"
	path += hold
	path += "/"
	path += "portfolio.json"
	fmt.Println("FindFile:", path)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Could not find the file")
		return "ERROR", "ERROR"
	}
	fmt.Println(string(file))

	data := Portfolio{}
	_ = json.Unmarshal([]byte(file), &data)

	return data.Information.Name, data.PortStatus.Status
}

// EditFile =
func EditFile(user string, status string) {
	path := "../Portfolios/"
	path += user
	path += "/"
	path += "portfolio.json"
	fmt.Println("PATH:", path)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Could not find the file")
	}
	fmt.Println(string(file))

	data := Portfolio{}
	_ = json.Unmarshal([]byte(file), &data)

	if !((status == "APPROVED" && data.PortStatus.Status == "APPROVED") || (status == "DENIED" && data.PortStatus.Status == "DENIED")) {
		//Delete old json file so new information isn't appended to existing file.
		cmd := exec.Command("rm", path)
		cmd.Run()

		//Create new json file of same name, all data is same except for modified status.
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return
		}
		defer f.Close()

		// Up to this point, "data" is a new json file that contains the exact information of the old json file.
		// If a choice 1 or 2 is made, Status will be changed before json file is remade.
		if status == "APPROVED" {
			data.PortStatus.Status = "APPROVED"
			fmt.Println("Status set to APPROVED")
		}

		if status == "DENIED" {
			data.PortStatus.Status = "DENIED"
			fmt.Println("Status set to DENIED")
		}

		b, _ := json.MarshalIndent(data, "", "    ") //new json file is remade with same name as old file.
		f.Write(b)
		f.Close()
	}
}

// ParsePayload = replace all encoded symbols to read payload information more easily
func ParsePayload(payload string) (slice []string) {
	var result string
	result = strings.ReplaceAll(payload, "%7B", " ")
	result = strings.ReplaceAll(result, "%22", " ")
	result = strings.ReplaceAll(result, "%3A", " ")
	result = strings.ReplaceAll(result, "%2C", " ")
	result = strings.ReplaceAll(result, "%5C", " ")
	result = strings.ReplaceAll(result, "%2F", " ")
	result = strings.ReplaceAll(result, "%7D", " ")
	result = strings.ReplaceAll(result, "%5D", " ")
	result = strings.ReplaceAll(result, "%5B", " ")
	result = strings.ReplaceAll(result, "%27", "'")
	result = strings.ReplaceAll(result, "%40", "@")

	slice = strings.Fields(result)
	return slice
}

// GetEmail = email value is at this specific location in the parsed payload slice
func GetEmail(slice []string) string {
	return slice[9]
}

// GetButton = button value is at this specific location in parsed payload slice
func GetButton(slice []string) string {
	return slice[5]
}

// ParseEmail = return email username before the "@" symbol
func ParseEmail(email string) (user string) {
	index := strings.Index(email, "@")
	if index == -1 {
		return user
	}

	user = email[0:index]
	return user
}
