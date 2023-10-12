package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"poc/messaging/mailgun/sender"
	"poc/shared/generic"
)

var (
	domain         = os.Getenv("MESSAGING_MAILGUN_DOMAIN")
	apiKey         = os.Getenv("MESSAGING_MAILGUN_API_KEY")
	emailSender    = os.Getenv("MESSAGING_EMAIL_SENDER")
	emailRecipient = os.Getenv("MESSAGING_EMAIL_RECIPIENT")
)

func main() {
	email := generic.Email{
		From:    emailSender,
		To:      emailRecipient,
		Subject: "Test",
		Message: "Hello, world!",
	}

	if err := sender.SendEmail(context.Background(), domain, apiKey, email); err != nil {
		log.Fatalf("failed to send e-mail: %v", err)
	}

	fmt.Printf("e-mail sent from %s to %s successfully!", email.From, email.To)
}
