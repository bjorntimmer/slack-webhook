package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Slack the client containing the webhook url
type Slack struct {
	WebhookURL string
}

// RequestBody the body of a slack message
type RequestBody struct {
	Text string `json:"text"`
}

// New create a new slack client including the webhook url
func New(webhookURL string) *Slack {
	return &Slack{
		WebhookURL: webhookURL,
	}
}

func (sl *Slack) Send(message string) error {
	slackBody, _ := json.Marshal(RequestBody{Text: message})
	req, err := http.NewRequest(http.MethodPost, sl.WebhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
