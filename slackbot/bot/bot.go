package bot

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

// User = holds User info
type User struct {
	ID          string
	AccountName string
	DisplayName string
	Email       string
}

// StartBot =
// Connect to RTM API and infinite loop to keep program running
// Read IncomingEvents message type for the switch statements
// It Connects, Responds, or Exits (when interrupted)
func StartBot() {
	// Connect to RTM API with Slack token
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	// Infinite Loop (break when interrupted)
Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event: ")
			switch event := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connected!")
				//fmt.Println("Connected! -> ", event.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Println("Message Received!")
				//fmt.Printf("%v\n", event)

				// Get bot info
				botInfo := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", botInfo.User.ID)

				// Get user info
				userInfo, err := api.GetUserInfo(event.User)
				if err != nil {
					fmt.Printf("%s\n", err)
					return
				}

				// If User DisplayName is not set use RealName
				displayName := userInfo.Profile.DisplayName
				if displayName == "" {
					displayName = userInfo.Profile.RealName
				}
				human := User{userInfo.ID, userInfo.Profile.RealName, displayName, userInfo.Profile.Email}

				// Check if message is from a User and is directed to the bot
				if event.User != botInfo.User.ID && strings.HasPrefix(event.Text, prefix) {
					Respond(rtm, event, prefix, human)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", event.Error())

			case *slack.InvalidAuthEvent:
				fmt.Println("Error -> Invalid Credentials")
				break Loop

			default:
				fmt.Printf("~")
			}
		}
		fmt.Println()
	}
}

// Respond = Bot responds using data from "data.go"
// Check if User message exists in map data and respond accordingly
// Append the User's display name to the messages
func Respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string, human User) {
	var response string
	text := TrimString(msg.Text, prefix)

	hiValue, hiBool := UserHi[text]
	byeValue, byeBool := UserBye[text]

	if hiBool {
		if hiValue == "rng" {
			response = BotHi[Random(0, len(BotHi))] + human.DisplayName
		} else {
			response = hiValue + human.DisplayName
		}
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if byeBool {
		if byeValue == "rng" {
			response = BotBye[Random(0, len(BotBye))] + human.DisplayName
		} else {
			response = byeValue + human.DisplayName
		}
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else {
		response = "Sorry I don't understand"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}

// TrimString = Remove the prefix from the msg
//			  = Convert msg to lower case
//			  = Remove all characters that are not letters or numbers
// (However this also removes whitespace)
func TrimString(msg string, prefix string) (result string) {
	result = strings.TrimPrefix(msg, prefix)
	result = strings.ToLower(result)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	result = reg.ReplaceAllString(result, "")

	fmt.Println("result:", result)
	return result
}

// Random = return a random int from [min, max)
func Random(min, max int) int {
	return min + rand.Intn(max-min)
}
