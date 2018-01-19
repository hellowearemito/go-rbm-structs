package rcs

import (
	b64 "encoding/base64"
	"encoding/json"

	"github.com/pkg/errors"
)

// Decode decodes the given data to the appropriate event or message struct.
func Decode(data []byte) (interface{}, error) {
	var event Event

	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal")
	}

	decodedData, err := b64.StdEncoding.DecodeString(event.Message.Data)
	if err != nil {
		return nil, errors.Wrap(err, "data base64 decode")
	}

	switch event.Message.Attributes.Type {
	case EventAttributeType:
		if event.Message.Attributes.EventType == nil {
			return nil, errors.New("Missing event type for event")
		}

		switch *event.Message.Attributes.EventType {
		case DeliveredEventType:
			var data EventMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				return nil, errors.Wrap(err, "json unmarshal delivered data.")
			}
			return DeliveredEventMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}, nil
		case ReadEventType:
			var data EventMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				return nil, errors.Wrap(err, "json unmarshal read data.")
			}
			return ReadEventMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}, nil
		case IsTypingEventType:
			var data EventMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				return nil, errors.Wrap(err, "json unmarshal read data.")
			}
			return IsTypingEventMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}, nil
		}

	case MessageAttributeType:
		if event.Message.Attributes.MessageType == nil {
			return nil, errors.New("Missing event type for event")
		}

		switch *event.Message.Attributes.MessageType {
		case TextMessageType:
			var data TextMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				return nil, errors.Wrap(err, "json unmarshal text data.")
			}
			return TextMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}, nil
		case UserFileMessageType:
			var data UserFileMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				return nil, errors.Wrap(err, "json unmarshal user file data.")
			}
			return UserFileMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}, nil
		case LocationMessageType:
			var data LocationMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				return nil, errors.Wrap(err, "json unmarshal location data")
			}
			return LocationMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}, nil
		case SuggestionResponseMessageType:
			var data SuggestionResponseMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				return nil, errors.Wrap(err, "json unmarshal suggestion response data")
			}
			return SuggestionResponseMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}, nil
		}
	}

	return nil, nil
}
