package pagerduty

import (
	"encoding/json"
	"io"
)

// V2WebhookPayload is a list of messages for a webhook.
type V2WebhookPayload struct {
	Messages []Message `json:"messages"`
}

// Message represents a single message in a v2 webhook payload.
type Message struct {
	ID         string     `json:"id"`
	Event      string     `json:"event"`
	CreatedOn  string     `json:"created_on"` // TODO: time.Time
	LogEntries []LogEntry `json:"log_entries"`
	Webhook    Webhook    `json:"webhook"`
	Incident   Incident   `json:"incident"`
}

// Webhook represents information about the webhook.
type Webhook struct {
	APIObject
	EndpointURL         string      `json:"endpoint_url"`
	Name                string      `json:"name"`
	Description         string      `json:"description"`
	WebhookObject       APIObject   `json:"webhook_object"`
	Config              interface{} `json:"config"` // Not sure what this is
	OutboundIntegration APIObject   `json:"outbound_integration"`
	AccountsAddon       []Addon     `json:"accounts_addon"` // Not sure what this is, I supposed that was a slice of addon
}

// DecodeV2Webhook decodes a webhook from an io.Reader.
func DecodeV2Webhook(r io.Reader) (V2WebhookPayload, error) {
	var payload V2WebhookPayload
	if err := json.NewDecoder(r).Decode(&payload); err != nil {
		return payload, err
	}
	return payload, nil
}
