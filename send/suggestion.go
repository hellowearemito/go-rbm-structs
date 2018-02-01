package send

import (
	"encoding/json"
	"errors"
)

// Suggestion represents the suggestion structure.
type Suggestion struct {
	SuggestedBase
	Action interface{}
}

// MarshalJSON for the suggestion structure.
func (s Suggestion) MarshalJSON() ([]byte, error) {
	var suggestion interface{}
	if s.Action == nil {
		suggestion = struct {
			Reply SuggestedBase `json:"reply,omitempty"`
		}{
			Reply: s.SuggestedBase,
		}
	} else {
		switch action := s.Action.(type) {
		case DialAction:
			suggestion = struct {
				Action dialActionWrapper `json:"action,omitempty"`
			}{
				Action: dialActionWrapper{
					SuggestedBase: s.SuggestedBase,
					DialAction:    action,
				},
			}
		case ViewLocationAction:
			suggestion = struct {
				Action viewLocationActionWrapper `json:"action,omitempty"`
			}{
				Action: viewLocationActionWrapper{
					SuggestedBase:      s.SuggestedBase,
					ViewLocationAction: action,
				},
			}
		case CreateCalendarEventAction:
			suggestion = struct {
				Action createCalendarEventActionWrapper `json:"action,omitempty"`
			}{
				Action: createCalendarEventActionWrapper{
					SuggestedBase:             s.SuggestedBase,
					CreateCalendarEventAction: action,
				},
			}
		case OpenURLAction:
			suggestion = struct {
				Action openURLActionWrapper `json:"action,omitempty"`
			}{
				Action: openURLActionWrapper{
					SuggestedBase: s.SuggestedBase,
					OpenURLAction: action,
				},
			}
		case ShareLocationAction:
			suggestion = struct {
				Action shareLocationActionWrapper `json:"action,omitempty"`
			}{
				Action: shareLocationActionWrapper{
					SuggestedBase:       s.SuggestedBase,
					ShareLocationAction: action,
				},
			}
		default:
			return nil, errors.New("invalid action type")
		}
	}

	return json.Marshal(suggestion)
}

// UnmarshalJSON for the suggestion structure.
func (s *Suggestion) UnmarshalJSON(b []byte) error {
	var suggestion suggestionUnmarshaller

	err := json.Unmarshal(b, &suggestion)
	if err != nil {
		return err
	}

	if suggestion.Reply != nil {
		s.SuggestedBase = *suggestion.Reply
		s.Action = nil
	} else if suggestion.Action != nil {
		s.SuggestedBase = suggestion.Action.SuggestedBase
		if suggestion.Action.DialAction != nil {
			s.Action = *suggestion.Action.DialAction
		} else if suggestion.Action.ViewLocationAction != nil {
			s.Action = *suggestion.Action.ViewLocationAction
		} else if suggestion.Action.CreateCalendarEventAction != nil {
			s.Action = *suggestion.Action.CreateCalendarEventAction
		} else if suggestion.Action.OpenURLAction != nil {
			s.Action = *suggestion.Action.OpenURLAction
		} else if suggestion.Action.ShareLocationAction != nil {
			s.Action = *suggestion.Action.ShareLocationAction
		}
	}
	return nil
}

// suggestionActionUnmarshaller represents the suggestion action unmarshaller structure.
type suggestionActionUnmarshaller struct {
	SuggestedBase
	DialAction                *DialAction
	ViewLocationAction        *ViewLocationAction
	CreateCalendarEventAction *CreateCalendarEventAction
	OpenURLAction             *OpenURLAction
	ShareLocationAction       *ShareLocationAction
}

// suggestionUnmarshaller represents the suggestion unmarshaller structure.
type suggestionUnmarshaller struct {
	Reply  *SuggestedBase
	Action *suggestionActionUnmarshaller
}

// SuggestedBase represents the base suggestion structure.
type SuggestedBase struct {
	Text         string `json:"text,omitempty"`
	PostbackData string `json:"postbackData,omitempty"`
}

// dialActionWrapper represnts the dial action wrapper structure.
type dialActionWrapper struct {
	SuggestedBase
	DialAction DialAction `json:"dialAction,omitempty"`
}

// viewLocationActionWrapper represents the view location wrapper structure.
type viewLocationActionWrapper struct {
	SuggestedBase
	ViewLocationAction ViewLocationAction `json:"viewLocationAction,omitempty"`
}

// createCalendarEventActionWrapper represents the create calendar event action wrapper structure.
type createCalendarEventActionWrapper struct {
	SuggestedBase
	CreateCalendarEventAction CreateCalendarEventAction `json:"createCalendarEventAction,omitempty"`
}

// openURLActionWrapper represents the open url action wrapper structure.
type openURLActionWrapper struct {
	SuggestedBase
	OpenURLAction OpenURLAction `json:"openUrlAction,omitempty"`
}

// shareLocationActionWrapper represents the share location action wrapper structure.
type shareLocationActionWrapper struct {
	SuggestedBase
	ShareLocationAction ShareLocationAction `json:"shareLocationAction,omitempty"`
}

// DialAction represents the dial action structure.
type DialAction struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// ViewLocationAction represents the view location action structure.
type ViewLocationAction struct {
	LatLong Location `json:"latLong,omitempty"`
	Label   string   `json:"label,omitempty"`
	Query   string   `json:"query,omitempty"`
}

// Location represents the location structure.
type Location struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

// CreateCalendarEventAction represents the create calendar event action structure.
type CreateCalendarEventAction struct {
	StartTime   string `json:"startTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// OpenURLAction represents the open url action structure.
type OpenURLAction struct {
	URL string `json:"url,omitempty"`
}

// ShareLocationAction represents the share location action structure.
type ShareLocationAction struct{}
