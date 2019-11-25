package main

/*
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

//Portfolio stores all of the information pulled from the portfolio form. This is stored as global variable.
type Portfolio struct {
	Information Information `json:"Information"`
	About       About       `json:"About"`
	Education   Education   `json:"Education"`
	Project     Project     `json:"Project"`
	Status      string
}

//Information stores all of the relevant info on the person creating the portfolio.
type Information struct {
	Name  string `json:"Name"`
	Title string `json:"Title"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}

//About stores the text the user enters in the "About Me" field.
type About struct {
	Aboutme string `json:"Aboutme"`
}

//Education stores the college the user attended and the degree they got, if any.
type Education struct {
	College string `json:"College"`
	Degree  string `json:"Degree"`
}

//Project stores information on the user's projects that they have completed, if any.
type Project struct {
	Name string `json:"Name"`
	Tech string `json:"Tech"`
	Desc string `json:"Desc"`
}

//userInput reads the filename inputted by the user.
var userInput string

func main() {

	fmt.Println("Type in filename of .json file (example.json)")
	fmt.Scanln(&userInput)

	//Reading, Unmarshaling file.
	file, _ := ioutil.ReadFile(userInput)
	data := Portfolio{}
	_ = json.Unmarshal([]byte(file), &data)

	//Terminal prompt to display status of portfolio, and prompt action
	fmt.Println("Portfolio Name is : ", data.Information.Name)
	fmt.Println("Portfolio Status is : ", data.Status)
	fmt.Println("What would you like to do with " + data.Information.Name + "'s portfolio? Press number to choose: ")
	fmt.Println()
	fmt.Println("	1: APPROVE portfolio")
	fmt.Println("	2: DENY portfolio")
	fmt.Println("	3: EXIT")
	fmt.Println()
	var choice int
	fmt.Printf("Enter choice: ")
	fmt.Scanln(&choice)

	//Delete old json file so new information isn't appended to existing file.
	remote3 := exec.Command("rm", userInput)
	remote3.Run()

	//Create new json file of same name, all data is same except for modified status.
	f, err := os.OpenFile(userInput, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	/* 	Up to this point, "data" is a new json file that contains the exact information of the
	old json file. If a choice 1 or 2 is made, Status will be changed before json file is remade.*/
/*
	switch choice {
	case 1:
		data.Status = "APPROVED"
		fmt.Println("Status set to APPROVED")
	case 2:
		data.Status = "DENIED"
		fmt.Println("Status set to DENIED")
	case 3:
		fmt.Println("Exiting . . .")
		b, _ := json.MarshalIndent(data, "", "    ") //new json file is remade with same name as old file.
		f.Write(b)
		f.Close()
		os.Exit(0)
	default:
	}
	b, _ := json.MarshalIndent(data, "", "    ") //new json file is remade with same name as old file.
	f.Write(b)
	f.Close()
}
*/
