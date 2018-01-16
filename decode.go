package rcs

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Decode decodes the given data to the appropriate event or message struct
func Decode(data []byte) (interface{}, error) {
	var raw rawMessage

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal")
	}

	switch raw.Attributes.Type {
	case EventAttributeType:
		if raw.Attributes.EventType == nil {
			return nil, errors.New("Missing event type for event")
		}
	case MessageAttributeType:
		if raw.Attributes.MessageType == nil {
			return nil, errors.New("Missing event type for event")
		}

		switch *raw.Attributes.MessageType {
		case TextMessageType:
			var data TextMessageData
			err = json.Unmarshal([]byte(raw.Data), &data)
			if err != nil {
				return nil, errors.Wrap(err, "json unmarshal text data")
			}
			return TextMessage{
				BaseMessage: raw.BaseMessage,
				Data:        data,
			}, nil
		}
	}

	return nil, nil
}

func akarmi() {
	message, _ := Decode([]byte(``))
	switch message := message.(type) {
	case TextMessage:
		fmt.Println(message.Data)
	}
}
