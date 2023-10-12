package sender

import (
	"github.com/slack-go/slack"
)

func NewTitle(title string) *slack.RichTextBlock {
	return slack.NewRichTextBlock("title",
		slack.NewRichTextSection(
			slack.NewRichTextSectionTextElement(title, &slack.RichTextSectionTextStyle{
				Bold: true,
			}),
		),
	)
}

func SendMessage(token string, channelID string, message string) error {
	client := slack.New(token)

	blocks := slack.MsgOptionBlocks(
		slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, message, false, false), nil, nil),
	)

	if _, _, err := client.PostMessage(channelID, blocks); err != nil {
		return err
	}

	return nil
}
