package main

import (
	"log"
	"os"

	"poc/messaging/slack/sender"
	"poc/shared/template"
)

var (
	token            = os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN")
	channelID        = os.Getenv("SLACK_BOT_TARGET_CHANNEL_ID")
	templateFilepath = "../../shared/assets/templates/message.md"
)

func main() {
	t, err := template.FromFile(templateFilepath)
	if err != nil {
		log.Fatalf("failed to create template: %v\n", err)
	}

	message := t.Build(template.Fields{
		"message_name": "some-message-name",
		"message_id":   "some-message-id",
	})

	if err := sender.SendMessage(token, channelID, message); err != nil {
		log.Fatalf("failed to send message: %v\n", err)
	}
}
