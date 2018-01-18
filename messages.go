package rcs

import (
	"time"
)

// DeliveredEventType contains the delivered event string.
const DeliveredEventType = "DELIVERED"

// IsTypingEventType contains the is typing event string.
const IsTypingEventType = "IS_TYPING"

// ReadEventType contains the read event string.
const ReadEventType = "READ"

// TextMessageType contains the type of text message string.
const TextMessageType = "TEXT"

// UserFileMessageType contains the type of user file message string.
const UserFileMessageType = "USER_FILE"

// LocationMessageType contains the type of location message string.
const LocationMessageType = "LOCATION"

// SuggestionResponseMessageType contains the suggestion response message string.
const SuggestionResponseMessageType = "SUGGESTION_RESPONSE"

// EventAttributeType contains the type of event attribute string.
const EventAttributeType = "event"

// MessageAttributeType contains the type of message attribute string.
const MessageAttributeType = "message"

// BaseMessage represents the base messages structure.
type BaseMessage struct {
	Attributes  Attributes
	MessageID   string
	PublishTime time.Time
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
	Data string
}

// BaseData represents the base data structure.
type BaseData struct {
	SenderPhoneNumber string
	MessageID         string
	SendTime          time.Time
}

// TextMessageData represents the text message data structure.
type TextMessageData struct {
	BaseData
	Text string
}

// TextMessage represents the text message structure.
type TextMessage struct {
	BaseMessage
	Data TextMessageData
}

// EventMessageData represents the event message data structure.
type EventMessageData struct {
	BaseData
	EventType string
	EventID   string
}

// UserFileMessageData represents the user file message structure.
type UserFileMessageData struct {
	BaseData
	UserFile UserFile
}

// UserFile represents the user file structure.
type UserFile struct {
	Payload   Payload
	Thumbnail *Thumbnail
}

// Payload represents the payload of user file structure.
type Payload struct {
	MimeType      string
	FileSizeBytes int
	FileName      string
	FileURI       string
}

// Thumbnail represents the thumbnail of user file structure.
type Thumbnail struct {
	MimeType      string
	FileSizeBytes int
	FileURI       string
}

// UserFileMessage represents the user file message structure.
type UserFileMessage struct {
	BaseMessage
	Data UserFileMessageData
}

// LocationMessage represents the location message structure.
type LocationMessage struct {
	BaseMessage
	Data LocationMessageData
}

// LocationMessageData represents the location data structure.
type LocationMessageData struct {
	BaseData
	Location Location
}

// Location represents the location field of location data structure.
type Location struct {
	Latitude  float32
	Longitude float32
}

// SuggestionResponseMessage represents the suggestion response message structure.
type SuggestionResponseMessage struct {
	BaseMessage
	Data SuggestionResponseMessageData
}

// SuggestionResponseMessageData represents the suggestion response data structure.
type SuggestionResponseMessageData struct {
	BaseData
	SuggestionResponse SuggestionResponse
}

// SuggestionResponse represents the suggestion response field of suggestion response data.
type SuggestionResponse struct {
	PostbackData string
	Text         string
}

// DeliveredEventMessage represents the delivered event message structure.
type DeliveredEventMessage struct {
	BaseMessage
	Data EventMessageData
}

// IsTypingEventMessage represents the is typing event message structure.
type IsTypingEventMessage struct {
	BaseMessage
	Data EventMessageData
}

// ReadEventMessage represents the read event message structure
type ReadEventMessage struct {
	BaseMessage
	Data EventMessageData
}
