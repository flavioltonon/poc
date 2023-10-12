package main

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"poc/messaging/slack/sender"
)

var (
	token     = os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN")
	channelID = os.Getenv("SLACK_BOT_TARGET_CHANNEL_ID")

	// Path to a mrkdwn template (not Markdown)
	templateFilepath = "./assets/templates/template.md"
)

func main() {
	t, err := template.ParseFiles(templateFilepath)
	if err != nil {
		log.Fatalf("failed to parse template: %v\n", err)
	}

	buffer := new(bytes.Buffer)

	if err := t.Execute(buffer, map[string]string{
		"message_name": "some-message-name",
		"message_id":   "some-message-id",
		"some_url":     "https://google.com",
	}); err != nil {
		log.Fatalf("failed to execute template: %v\n", err)
	}

	if err := sender.SendMessage(token, channelID, buffer.String()); err != nil {
		log.Fatalf("failed to send message: %v\n", err)
	}
}
