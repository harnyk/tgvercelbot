# tgvercelbot

Telegram Webhook Bot Handler for Vercel

## What does it do

`tgvercelbot` is a Telegram Webhook Bot Handler for Vercel.

It allows you to easely host your Telegram Bot with Vercel.

## What doesn't it do

**`tgvercelbot` doesn't setup a webhook for your bot**, you need to do that yourself, or use a [tgvercel](https://github.com/harnyk/tgvercel) tool, which is specially designed to automate the Telegram webhooks setup for Vercel.

## How to use

Suppose, you have set up the Telegram webhook for your bot so that it points to the `api/tg/webhook?secret=YOUR_SECRET` endpoint in your project.

Also let's suppose that your Vercel project has the following environment variables already set:

-   `TELEGRAM_TOKEN` - your bot token
-   `TELEGRAM_WEBHOOK_SECRET` - the YOUR_SECRET part from the webhook URL. Can be just a random string.

BTW, you can easily configure these environment variables using [tgvercel](https://github.com/harnyk/tgvercel) tool (the command is `tgvercel init --target=preview --telegram-token=TELEGRAM_BOT_TOKEN`). But also you can do it manually.

So you can build the handler like the following:

```go
package handler

import (
	"net/http"

	"github.com/harnyk/tgvercelbot"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var tgv = tgvercelbot.New(tgvercelbot.DefaultOptions())

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	tgv.HandleWebhook(r, func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
        if update.Message != nil {
            // do something, for example, echo:
            bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
        }
    })
}
```

The `tgvercelbot` will pick up both environment variables from your project and use it to initialize the bot API client, as well as to check the authenticity of the request comparing the TELEGRAM_WEBHOOK_SECRET value with the `secret` query parameter of the webhook request.

After you deploy your bot to Vercel, don't forget to point the Telegram webhook URL to `api/tg/webhook` endpoint for your specific deployment. You can do it manually or using [tgvercel](https://github.com/harnyk/tgvercel) tool (the command is `tgvercel hook DEPLOYMENT_ID_OR_URL`).

## Running locally

It is not possible to run your webhook-based endpoint locally. But you can create the executable which would be able to run your bot using polling mode:

```go
package main

import (
	"log"

	"github.com/harnyk/tgvercelbot"
	"github.com/joho/godotenv"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func main() {
	log.Println("Starting bot locally...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	err = tgvercelbot.RunLocal(env.MustGet("TELEGRAM_TOKEN"), func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
        if update.Message != nil {
            // do something, for example, echo:
            bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
        }
    })
	if err != nil {
		log.Fatalf("failed to run locally: %v", err)
	}

}
```

The second argument of `RunLocal` function is a callback function which must be the exactly same as the one in `WebhookHandler` function above.

You must specify `TELEGRAM_TOKEN` environment variable in your `.env` file.

**Heads up**: running `RunLocal` function will remove the existing webhook for your bot. That's why you should never use the same Telegram token as in preview or production for the development environment.
