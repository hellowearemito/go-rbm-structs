package rcs

// Event represents the event structure.
type Event struct {
	Message      rawMessage   `json:"message"`
	Subscription Subscription `json:"subscription"`
}
