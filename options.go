package tgvercelbot

import (
	"fmt"
)

type Options struct {
	TelegramTokenEnvName         string
	TelegramWebhookSecretEnvName string
}

func DefaultOptions() Options {
	return Options{
		TelegramTokenEnvName:         "TELEGRAM_TOKEN",
		TelegramWebhookSecretEnvName: "TELEGRAM_WEBHOOK_SECRET",
	}
}

func (o *Options) validate() error {
	if o.TelegramTokenEnvName == "" {
		return fmt.Errorf("TelegramTokenEnvName must be set")
	}

	if o.TelegramWebhookSecretEnvName == "" {
		return fmt.Errorf("TelegramWebhookSecretEnvName must be set")
	}

	return nil
}
