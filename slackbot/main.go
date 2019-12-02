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

	// Open Postgres database
	//database.StartDB()

	http.HandleFunc("/", home)
	//http.ListenAndServe(":6666", nil)
	go func() {
		err := http.ListenAndServe(":6666", nil)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	// Start the bot
	bot.StartBot()

	//database.FindFile("davidychang@outlook.com")
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
		//fmt.Println(string(result))
		slice := database.ParsePayload(string(result))
		user := database.ParseEmail(database.GetEmail(slice))
		button := database.GetButton(slice)

		//fmt.Println(database.GetButton(slice))
		//fmt.Println(database.GetEmail(slice))
		//fmt.Println(database.ParseEmail(database.GetEmail(slice)))
		//fmt.Println(database.GetValue(string(result)))
		//fmt.Println(database.ParseEmail(database.GetValue(string(result))))
		database.EditFile(user, button)
		w.Write([]byte(database.GetButton(slice)))
		//w.Write([]byte("Received a POST request\n"))

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}
