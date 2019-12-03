package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/190930-UTA-CW-Go/project3/slackbot/database"
)

func main() {
	// Set RNG Seed
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", home)
	go func() {
		err := http.ListenAndServe(":6666", nil)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	// Start the bot
	//bot.StartBot()

	//getstatus("Volcano")
	FindFile("davidychang")
}

const amazon = "ec2-user@ec2-18-188-174-65.us-east-2.compute.amazonaws.com" //Amazon aws hostname@IPaddress.
const key = "../rego.pem"                                                   //This is security key used to login/SSH

// FindFile =
func FindFile(username string) {
	cmd := exec.Command("scp", "-i", key, amazon+":/home/ec2-user/Portfolios/"+username+"/"+username+".json", ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + "; " + string(output))
		return
	}
	fmt.Println(fmt.Sprint(err) + ": " + string(output))

	fmt.Println("Downloading file....")

}

func potato(username string) {
	// Parse out "Name"
	cmd1 := " cat Portfolios/" + username + "/" + username + ".json" + " | egrep 'Name'"
	output1, err1 := exec.Command("bash", "-c", ("ssh -i " + key + " " + amazon + cmd1)).Output()
	if err1 != nil {
		fmt.Println("ERROR")
	}
	//fmt.Println(string(output1))
	//slice1 := strings.Fields(string(output1))
	slice1 := strings.Split(string(output1), "\n")
	hold := slice1[0]
	fmt.Println(hold[8:])
	/*
		file, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("Could not find the file")
			return "ERROR", "ERROR"
		}
		fmt.Println(string(file))

		data := Portfolio{}
		_ = json.Unmarshal([]byte(file), &data)

		return data.Information.Name, data.PortStatus.Status
	*/
}

func getstatus(username string) string {
	output, err := exec.Command("bash", "-c", ("ssh -i " + key + " " + amazon + " stat Portfolios/" + username + " | egrep 'Modify'")).Output()
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Println(string(output))
	slice := strings.Fields(string(output))

	dateTime := slice[1] + " " + (slice[2])[:5]
	fmt.Println("dateTime:", dateTime)

	t, _ := time.Parse("2006-01-02 15:04", dateTime)
	result := t.Format("January 2, 2006 at 3:04 PM")
	fmt.Println(result)
	return result
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		w.Write([]byte("Received a GET request\n"))

	case "POST":
		result, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		slice := database.ParsePayload(string(result))
		user := database.ParseEmail(database.GetEmail(slice))
		button := database.GetButton(slice)
		database.EditFile(user, button)
		w.Write([]byte(database.GetButton(slice)))

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}
