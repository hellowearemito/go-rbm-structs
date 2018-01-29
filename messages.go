package rcs

import (
	"time"
)

const (
	// DeliveredEventType contains the delivered event string.
	DeliveredEventType = "DELIVERED"
	// IsTypingEventType contains the is typing event string.
	IsTypingEventType = "IS_TYPING"
	// ReadEventType contains the read event string.
	ReadEventType = "READ"
)

const (
	// TextMessageType contains the type of text message string.
	TextMessageType = "TEXT"
	// UserFileMessageType contains the type of user file message string.
	UserFileMessageType = "USER_FILE"
	// LocationMessageType contains the type of location message string.
	LocationMessageType = "LOCATION"
	// SuggestionResponseMessageType contains the suggestion response message string.
	SuggestionResponseMessageType = "SUGGESTION_RESPONSE"
)

const (
	// EventAttributeType contains the type of event attribute string.
	EventAttributeType = "event"
	// MessageAttributeType contains the type of message attribute string.
	MessageAttributeType = "message"
)

// BaseMessage represents the base messages structure.
type BaseMessage struct {
	Attributes  Attributes `json:"attributes"`
	MessageID   string     `json:"messageId"`
	PublishTime time.Time  `json:"publishTime"`
}

// Attributes represents the attributes structure.
type Attributes struct {
	Type        string
	MessageType *string `json:"message_type"`
	EventType   *string `json:"event_type"`
}

// rawMessage represents the raw messages structure.
type rawMessage struct {
	BaseMessage
	Data string `json:"data"`
}

// BaseData represents the base data structure.
type BaseData struct {
	SenderPhoneNumber string    `json:"senderPhoneNumber"`
	MessageID         string    `json:"messageId"`
	SendTime          time.Time `json:"sendTime"`
}

// TextMessageData represents the text message data structure.
type TextMessageData struct {
	BaseData
	Text string `json:"text"`
}

// TextMessage represents the text message structure.
type TextMessage struct {
	BaseMessage
	Data TextMessageData `json:"data"`
}

// EventMessageData represents the event message data structure.
type EventMessageData struct {
	BaseData
	EventType string `json:"eventType"`
	EventID   string `json:"eventId"`
}

// UserFileMessageData represents the user file message structure.
type UserFileMessageData struct {
	BaseData
	UserFile UserFile `json:"userFile"`
}

// UserFile represents the user file structure.
type UserFile struct {
	Payload   Payload    `json:"payload"`
	Thumbnail *Thumbnail `json:"thumbnail"`
}

// Payload represents the payload of user file structure.
type Payload struct {
	MimeType      string `json:"mimeType"`
	FileSizeBytes int    `json:"fileSizeBytes"`
	FileName      string `json:"fileName"`
	FileURI       string `json:"fileUri"`
}

// Thumbnail represents the thumbnail of user file structure.
type Thumbnail struct {
	MimeType      string `json:"mimeType"`
	FileSizeBytes int    `json:"fileSizeBytes"`
	FileURI       string `json:"fileUri"`
}

// UserFileMessage represents the user file message structure.
type UserFileMessage struct {
	BaseMessage
	Data UserFileMessageData `json:"data"`
}

// LocationMessage represents the location message structure.
type LocationMessage struct {
	BaseMessage
	Data LocationMessageData `json:"data"`
}

// LocationMessageData represents the location data structure.
type LocationMessageData struct {
	BaseData
	Location Location `json:"location"`
}

// Location represents the location field of location data structure.
type Location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// SuggestionResponseMessage represents the suggestion response message structure.
type SuggestionResponseMessage struct {
	BaseMessage
	Data SuggestionResponseMessageData `json:"data"`
}

// SuggestionResponseMessageData represents the suggestion response data structure.
type SuggestionResponseMessageData struct {
	BaseData
	SuggestionResponse SuggestionResponse `json:"suggestionResponse"`
}

// SuggestionResponse represents the suggestion response field of suggestion response data.
type SuggestionResponse struct {
	PostbackData string
	Text         string `json:"text"`
}

// DeliveredEventMessage represents the delivered event message structure.
type DeliveredEventMessage struct {
	BaseMessage
	Data EventMessageData `json:"data"`
}

// IsTypingEventMessage represents the is typing event message structure.
type IsTypingEventMessage struct {
	BaseMessage
	Data EventMessageData `json:"data"`
}

// ReadEventMessage represents the read event message structure
type ReadEventMessage struct {
	BaseMessage
	Data EventMessageData `json:"data"`
}
