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

				// Get bot info
				botInfo := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", botInfo.User.ID)

				// Get user info
				userInfo, err := api.GetUserInfo(event.User)
				if err != nil {
					fmt.Println("USER INFO ERROR:", err)
					goto Loop
				}

				// If User DisplayName is not set use RealName
				displayname := userInfo.Profile.DisplayName
				if displayname == "" {
					displayname = userInfo.Profile.RealName
				}

				// Check message is from a User and it's directed to the bot
				if event.User != botInfo.User.ID && strings.HasPrefix(event.Text, prefix) {
					// Extract username from email
					user := database.ParseEmail(userInfo.Profile.Email)
					Respond(rtm, event, api, prefix, user, displayname, channel)
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
func Respond(rtm *slack.RTM, msg *slack.MessageEvent, api *slack.Client, prefix string, user string, displayname string, channel string) {
	var response string

	// To trigger "review"
	// Ex) @gopher review username
	if strings.Contains(msg.Text, "review") {
		// Check user triggering review has admin rights
		if database.CheckAdmin(user) == false {
			response = "Need Administrator Permission."
			rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))

		} else {
			// Extract "username" from message
			slice := strings.Fields(msg.Text)
			if len(slice) == 3 {
				flag, name, status := database.FindFile(slice[2])
				if flag == false && name == "ERROR" && status == "ERROR" {
					response = "Sorry could not find folder."
					rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
				} else if flag == true && name == "" && status == "" {
					response = slice[2] + " has not started their portfolio yet."
					rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
				} else {
					ButtonMenu(rtm, api, channel, slice[2], name, status)
				}
			} else {
				response = "Sorry could not recognize command."
				rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
			}
		}
	} else {

		text := TrimString(msg.Text, prefix)

		// Checks which map is triggered
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
			if database.NewFolder(user) == true {
				if portValue == "rng" {
					response = BotPortfolio[Random(0, len(BotPortfolio))]
				} else {
					response = portValue + displayname
				}
			} else {
				response = "Sorry you might already have a folder."
			}
		} else if statusBool {
			statusFlag, statusValue := database.GetStatus(user)
			if statusFlag == true {
				response = "Last Updated: " + statusValue
			} else {
				response = "Sorry I could not find your portfolio."
			}
		} else {
			response = "Sorry I don't understand."
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

// ButtonMenu = sets up the look of the button and sends it
func ButtonMenu(rtm *slack.RTM, api *slack.Client, channel string, username string, name string, status string) {
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
				Value: username,
			},
			slack.AttachmentAction{
				Name:  "DENIED",
				Text:  "Deny",
				Type:  "button",
				Value: username,
				Style: "danger",
			},
		},
	}

	message := slack.MsgOptionAttachments(attachment)
	channelID, timestamp, err := api.PostMessage(channel, slack.MsgOptionText("", false), message)

	////////////////////////////////////////////////////////
	// Changed ListenAndServe to closure goroutine so these messages won't show anymore
	if err != nil {
		fmt.Println("BUTTON ERROR: Failed to send message")
	}
	fmt.Printf("BUTTON SUCCESS: Message sent to channel %s at %s", channelID, timestamp)
	///////////////////////////////////////////////////////

}
