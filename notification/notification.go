package notification

import (
	"fmt"
	"time"

	holiday "github.com/najeira/jpholiday"
	"github.com/slack-go/slack"
)

// Send ...
func Send(url string) error {
	today := time.Now()

	if isHolyday(today) {
		fmt.Println("今日は" + holiday.Name(today) + "!")
		return nil
	}

	text := "<!here>\n今日は平日!"

	msg := slack.WebhookMessage{
		Text: text,
	}

	return slack.PostWebhook(url, &msg)
}

func isHolyday(t time.Time) bool {
	return holiday.Name(t) != ""
}
