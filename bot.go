package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"net/url"
	"os"
)

func main() {
	slackToken := os.Getenv("SLACK_TOKEN")
	if slackToken == "" {
		fmt.Printf("SLACK_TOKEN env var not defined")
		return
	}
	api := slack.New(slackToken)
	rtm := api.NewRTM(slack.RTMOptionConnParams(url.Values{
		"batch_presence_aware": {"1"},
	}))
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			message := parse(ev.Msg.Text)
			fmt.Printf("Parsed message: %v\n", message)
			sendMessage(api, message)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		}
	}
}

func sendMessage(api *slack.Client, message Message) {
	destination, timestamp, err := api.PostMessage(message.Destination, slack.MsgOptionAsUser(true), slack.MsgOptionText(message.Text, false))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Message successfully sent to destination %s at %s\n", destination, timestamp)
	}
}
