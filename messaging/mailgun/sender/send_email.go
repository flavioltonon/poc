package sender

import (
	"context"
	"poc/shared/generic"
	"time"

	mailgun "github.com/mailgun/mailgun-go/v4"
)

func SendEmail(ctx context.Context, domain, apiKey string, email generic.Email) error {
	client := mailgun.NewMailgun(domain, apiKey)

	message := client.NewMessage(email.From, email.Subject, email.Message, email.To)

	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if _, _, err := client.Send(cctx, message); err != nil {
		return err
	}

	return nil
}
