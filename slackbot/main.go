package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/190930-UTA-CW-Go/project3/slackbot/bot"
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
	bot.StartBot()
}

// Handler for Slack API button
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

		// 1) Parse encoded JSON payload
		// 2) Get user email name
		// 3) Get button value ("APPROVED" / "DENIED")
		// 4) Edit the JSON file accordingly and upload to AWS machine
		// 5) Write to Slack chat depending on result
		slice := database.ParsePayload(string(result))
		user := database.GetEmail(slice)
		button := database.GetButton(slice)
		database.EditFile(user, button)

		if button == "APPROVED" {
			w.Write([]byte(user + "'s portfolio has been approved."))
		}
		if button == "DENIED" {
			w.Write([]byte(user + "'s portfolio has been denied."))
		}

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}
