package rcs

import (
	b64 "encoding/base64"
	"encoding/json"

	"github.com/pkg/errors"
)

// Decode decodes the given data to the appropriate event or message struct.
func Decode(data []byte) (subscription string, returnData interface{}, err error) {
	var event Event

	err = json.Unmarshal(data, &event)
	if err != nil {
		err = errors.Wrap(err, "json unmarshal")
		return
	}

	subscription = string(event.Subscription)

	decodedData, err := b64.StdEncoding.DecodeString(event.Message.Data)
	if err != nil {
		err = errors.Wrap(err, "data base64 decode")
		return
	}

	switch event.Message.Attributes.Type {
	case EventAttributeType:
		if event.Message.Attributes.EventType == nil {
			err = errors.New("Missing event type for event")
			return
		}

		switch *event.Message.Attributes.EventType {
		case DeliveredEventType:
			var data EventMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				err = errors.Wrap(err, "json unmarshal delivered data.")
				return
			}
			returnData = DeliveredEventMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}

			return
		case ReadEventType:
			var data EventMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				err = errors.Wrap(err, "json unmarshal read data.")
				return
			}
			returnData = ReadEventMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}

			return
		case IsTypingEventType:
			var data EventMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				err = errors.Wrap(err, "json unmarshal read data.")
				return
			}
			returnData = IsTypingEventMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}

			return
		}

	case MessageAttributeType:
		if event.Message.Attributes.MessageType == nil {
			err = errors.New("Missing event type for event")
			return
		}

		switch *event.Message.Attributes.MessageType {
		case TextMessageType:
			var data TextMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				err = errors.Wrap(err, "json unmarshal text data.")
				return
			}
			returnData = TextMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}

			return
		case UserFileMessageType:
			var data UserFileMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				err = errors.Wrap(err, "json unmarshal user file data.")
				return
			}
			returnData = UserFileMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}

			return
		case LocationMessageType:
			var data LocationMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				err = errors.Wrap(err, "json unmarshal location data")
				return
			}
			returnData = LocationMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}

			return
		case SuggestionResponseMessageType:
			var data SuggestionResponseMessageData
			err = json.Unmarshal(decodedData, &data)
			if err != nil {
				err = errors.Wrap(err, "json unmarshal suggestion response data")
				return
			}
			returnData = SuggestionResponseMessage{
				BaseMessage: event.Message.BaseMessage,
				Data:        data,
			}

			return
		}
	}

	return
}
