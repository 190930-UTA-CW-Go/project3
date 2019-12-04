RevatureGo-Slackbot
========

- Starts the bot.
- Gets user input.
- Bot uses data.go to give a response based upon user input.
- Makes folders when asked.
- Starts application when asked.
- Checks for folders and files and checks their status.
- Upload and Download files.

# Tutorial + Documentation
``` bash
1) Getting Started
   https://api.slack.com/bot-users#setup-events-api

2) Set "SLACK_API" environment variable
   export SLACK_API="______________"

3) You dont need to enable "Events API" for this program
   Instead enable "Interactive Components" in Slack API configuration 
   https://api.slack.com/tutorials/tunneling-with-ngrok

   sudo snap install ngrok
   ngrok http 6666
   (this creates a temporary session URL will need to update refresh after it expires)

4) Using Slack Go package "nlopes/slack"
   go get github.com/nlopes/slack

5) Invite bot to Slack chat
   Example) "/invite @gopher"

6) Run program while in Slacks folder
   go run main.go

7) Talk to bot in Slack
   Example) @gopher hi
```


# Commands
```bash
For command list can look at "slackbot/bot/data.go"
Commands are parsed of punctuations
Examples:
"Hellos"
@gopher hi
@gopher hello
@gopher hi!!!! (will work because punctuations are parsed)

"Goodbyes"
@gopher bye
@gopher goodbye

"Portfolio Creation"
// will automatically create and name Portfolio folder after Slack email username on AWS machine
@gopher create portfolio
@gopher start portfolio

"Portfolio Status"
// checks for Portfolio folder named after Slack email username
// then returns date and time Portfolio folder was last modified
@gopher portfolio status
@gopher status portfolio

"Review"
// only usernames listed in "admin.txt" will be able to trigger this command
// username is based on Slack email (without the @)
// will create button to respond to and update the json file on AWS machine accordingly
@gopher review username
```
