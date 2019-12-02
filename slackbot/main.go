package main

import (
	"math/rand"
	"time"

	"github.com/190930-UTA-CW-Go/project3/slackbot/bot"
)

func main() {
	// Set RNG Seed
	rand.Seed(time.Now().UTC().UnixNano())

	// Open Postgres database
	//database.StartDB()

	// Start the bot
	bot.StartBot()

	//database.FindFile("davidychang@outlook.com")
}
