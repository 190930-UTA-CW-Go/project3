package bot

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"github.com/190930-UTA-CW-Go/project3/slackbot/database"
	"github.com/nlopes/slack"
)

// StartBot =
// Connect to RTM API and infinite loop to keep program running
// Read "IncomingEvents" message type for the switch statements
// It Connects, Responds, or Exits (when interrupted)
func StartBot() {
	// Connect to RTM API with Slack token
	//token := os.Getenv("SLACK_TOKEN")
	token := os.Getenv("SLACK_API")
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

				channel := event.Channel
				fmt.Println("CHANNEL:", channel)
				/*
					idk, _ := rtm.GetIMChannels()
					fmt.Println(idk)
				*/

				// Get bot info
				botInfo := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", botInfo.User.ID)

				// Get user info
				userInfo, err := api.GetUserInfo(event.User)
				if err != nil {
					fmt.Println("USER INFO ERROR:", err)
					goto Loop
					//return
				}

				// If User DisplayName is not set use RealName
				displayname := userInfo.Profile.DisplayName
				if displayname == "" {
					displayname = userInfo.Profile.RealName
				}

				// Check message is from a User and it's directed to the bot
				if event.User != botInfo.User.ID && strings.HasPrefix(event.Text, prefix) {
					//////////////////////////////////////////////////
					// Instead of "userInfo.ID"
					id := database.ParseEmail(userInfo.Profile.Email)
					//////////////////////////////////////////////////
					database.Insert(id, userInfo.Profile.RealName, displayname, userInfo.Profile.Email)
					Respond(rtm, event, api, prefix, id, displayname, channel)
				}

			case *slack.RTMError:
				fmt.Printf("RTM ERROR: %s\n", event.Error())

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
func Respond(rtm *slack.RTM, msg *slack.MessageEvent, api *slack.Client, prefix string, id string, displayname string, channel string) {
	var response string
	if strings.Contains(msg.Text, "review") {
		//Review(rtm, msg, api)
		slice := strings.Fields(msg.Text)
		if len(slice) == 3 {
			index := strings.Index(slice[2], "|")
			email := (slice[2][index+1 : len(slice[2])-1])
			fmt.Println(email)
			name, status := database.FindFile(email)
			if name == "ERROR" && status == "ERROR" {
				response = "Sorry could not find folder"
				rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
			} else {
				ButtonMenu(rtm, api, channel, email, name, status)
			}
			//fmt.Println(name + ": " + status)
			//ButtonMenu(rtm, api, channel, name, status)

		}
	} else {

		text := TrimString(msg.Text, prefix)

		hiValue, hiBool := UserHi[text]
		byeValue, byeBool := UserBye[text]
		portValue, portBool := Portfolio[text]
		_, statusBool := Status[text]

		if hiBool {
			if hiValue == "rng" {
				response = BotHi[Random(0, len(BotHi))] + displayname
			} else {
				response = hiValue + displayname
			}
		} else if byeBool {
			if byeValue == "rng" {
				response = BotBye[Random(0, len(BotBye))] + displayname
			} else {
				response = byeValue + displayname
			}
		} else if portBool {
			if database.NewFolder(id) == true {
				if portValue == "rng" {
					response = BotPortfolio[Random(0, len(BotPortfolio))]
				} else {
					response = portValue + displayname
				}
			} else {
				response = "Sorry you might already have a folder"
			}
		} else if statusBool {
			if database.CompareStatus(id) == true {
				response = "Your portfolio is updated!"
			} else {
				response = "Your portfolio is not updated yet"
			}
		} else {
			response = "Sorry I don't understand"
		}
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}

// TrimString =
// Remove the prefix from the msg
// Convert msg to lower case
// Remove all characters that are not letters or numbers
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

// ButtonMenu =
//func ButtonMenu(rtm *slack.RTM, api *slack.Client, channel string, name string, status string) {
func ButtonMenu(rtm *slack.RTM, api *slack.Client, channel string, email string, name string, status string) {
	var pretext string = name + ": " + status
	attachment := slack.Attachment{
		Pretext:    pretext,
		Fallback:   "We don't currently support your client",
		CallbackID: "accept_or_reject",
		Color:      "#3AA3E3",
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "APPROVED",
				Text:  "Approve",
				Type:  "button",
				Value: email,
			},
			slack.AttachmentAction{
				Name:  "DENIED",
				Text:  "Deny",
				Type:  "button",
				Value: email,
				Style: "danger",
			},
		},
	}

	message := slack.MsgOptionAttachments(attachment)
	channelID, timestamp, err := api.PostMessage(channel, slack.MsgOptionText("", false), message)
	//channelID, timestamp, err := rtm.PostMessage(channel, slack.MsgOptionText("", false), message)
	if err != nil {
		fmt.Printf("Could not send message: %v", err)
		fmt.Println()
	}
	fmt.Printf("Message with buttons sucessfully sent to channel %s at %s", channelID, timestamp)
}
