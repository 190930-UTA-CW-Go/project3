package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

func main() {
	StartBot()
}

// StartBot = Connect to RTM API and infinite loop while program is running
//			=
func StartBot() {
	// Connect to RTM API with Slack token
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

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
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				// Check it's a User communicating with the bot
				// Check prefix exists in event.Text to identify it's the bot
				if event.User != info.User.ID && strings.HasPrefix(event.Text, prefix) {
					Respond(rtm, event, prefix)
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

// Respond = Bot respond using hardcoded input and output
func Respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	var response string
	text := TrimString(msg.Text, prefix)

	hiResponse := map[string]bool{
		"hi":    true,
		"hello": true,
		"hey":   true,
		"yo":    true,
	}

	byeResponse := map[string]bool{
		"goodbye": true,
		"bye":     true,
		"peace":   true,
	}

	if hiResponse[text] {
		response = "What's up?"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if byeResponse[text] {
		response = "Goodbye"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else {
		response = "Sorry I don't understand"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}

// TrimString = Remove prefix from msg
//			  = Lower case the msg
//			  = Remove characters that are not numbers or letters
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
