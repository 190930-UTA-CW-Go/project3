package main

import (
	"math/rand"
	"time"

	"github.com/190930-UTA-CW-Go/project3/slackbot/bot"
)

func main() {
	// Set RNG Seed
	rand.Seed(time.Now().UTC().UnixNano())

	// Start the bot
	bot.StartBot()
}
