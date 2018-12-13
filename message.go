package main

import "strings"

type Message struct {
	Destination string
	Text        string
}

func parse(text string) Message {
	message := new(Message)
	words := strings.Fields(text)
	message.Destination = parseDestination(words[0])
	message.Text = strings.Join(words[1:], " ")
	return *message
}

func parseDestination(encodedDestination string) string {
	cleanPrefix := strings.TrimPrefix(encodedDestination, "<#")
	cleanSuffix := strings.Split(cleanPrefix, "|")[0]
	return cleanSuffix
}
