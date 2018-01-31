package send

// Event represents the event message structure.
type Event struct {
	Name      string `json:"name"`
	EventType string `json:"eventType"`
	MessageID string `json:"messageId"`
	SendTime  string `json:"sendTime"`
}
