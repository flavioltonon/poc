# Slack messenger

Sends a mrkdwn message to a Slack channel.

## Inputs

- `SLACK_BOT_USER_OAUTH_TOKEN`: access token required for connecting to Slack API. Requires an Slack app to be created at [Your apps](https://api.slack.com/apps).
- `SLACK_BOT_TARGET_CHANNEL_ID`: ID (not the name!) of the channel the message should be sent to. It can be found by checking the details of a channel.

## References

- Formatting text for app surfaces: https://api.slack.com/reference/surfaces/formatting
- Creating an Slack app - Quickstart: https://api.slack.com/start/quickstart