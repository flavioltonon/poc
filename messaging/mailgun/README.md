# Mailgun E-mail API

## Inputs

- `MESSAGING_MAILGUN_DOMAIN`: Mailgun domain. Can be found [here](https://app.mailgun.com/app/sending/domains)
- `MESSAGING_MAILGUN_API_KEY`: key required to make API calls to Mailgun. Can be created [here](https://app.mailgun.com/settings/api_security)
- `MESSAGING_EMAIL_SENDER`: e-mail address of the sender of the message.
- `MESSAGING_EMAIL_RECIPIENT`: e-mail address of the recipient of the message. You may need to authorize specific recipients, if you are using a sandbox domain.

## Outputs

E-mail sent to the `MESSAGING_EMAIL_RECIPIENT`.
