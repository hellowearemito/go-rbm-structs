package send

// Event represents the event message structure.
type Event struct {
	EventType string `json:"eventType"`
	MessageID string `json:"messageId,omitempty"`
}
