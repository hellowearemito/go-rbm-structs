package send

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

// repleyJSON is a suggestion reply example.
const replyJSON = `{"reply":{"text":"text","postbackData":"postback data"}}`

// dialActionJson is a dial action suggestion example.
const dialActionJSON = `{"action":{"text":"text","postbackData":"postback data","dialAction":{"phoneNumber":"xyz"}}}`

// viewLocationActionJSON is a view location action suggestion example.
const viewLocationActionJSON = `{"action":{"text":"text","postbackData":"postback data","viewLocationAction":{"latLong":{"latitude":4.12,"longitude":2.23},"label":"label","query":"query"}}}`

// timeFormat is the format of the start and end time.
const timeFormat = "2014-10-02T15:01:23.045123456Z"

// start is an example for the calendar action start time.
var start = time.Now().Format(timeFormat)

// end is a example for the calendar action end time.
var end = time.Now().Format(timeFormat)

// createCalendarEventAction is a create calenendar event action suggestion example.
var createCalendarEventActionJSON = `{"action":{"text":"text","postbackData":"postback data","createCalendarEventAction":{"startTime":"` + start + `","endTime":"` + end + `","title":"title","description":"description"}}}`

// openURLActionJSON is an example for the open url action.
var openURLActionJSON = `{"action":{"text":"text","postbackData":"postback data","openUrlAction":{"url":"https://www.google.com"}}}`

func TestSuggestion(t *testing.T) {
	Convey("Suggestion", t, func() {
		Convey("Reply", func() {
			var suggestion Suggestion
			suggestion.SuggestedBase = SuggestedBase{
				Text:         "text",
				PostbackData: "postback data",
			}

			b, err := json.Marshal(suggestion)
			So(err, ShouldBeNil)
			So(bytes.Compare(b, []byte(replyJSON)), ShouldEqual, 0)

			err = json.Unmarshal(b, &suggestion)
			So(err, ShouldBeNil)
			So(suggestion.SuggestedBase.Text, ShouldEqual, "text")
			So(suggestion.SuggestedBase.PostbackData, ShouldEqual, "postback data")
			So(suggestion.Action, ShouldBeNil)
		})

		Convey("Dail Action", func() {
			var suggestion Suggestion
			suggestion.SuggestedBase = SuggestedBase{
				Text:         "text",
				PostbackData: "postback data",
			}
			suggestion.Action = DialAction{
				PhoneNumber: "xyz",
			}

			b, err := json.Marshal(suggestion)
			So(err, ShouldBeNil)
			So(bytes.Compare(b, []byte(dialActionJSON)), ShouldEqual, 0)

			err = json.Unmarshal(b, &suggestion)
			So(err, ShouldBeNil)
			So(suggestion.SuggestedBase.Text, ShouldEqual, "text")
			So(suggestion.SuggestedBase.PostbackData, ShouldEqual, "postback data")
			So(suggestion.Action, ShouldNotBeNil)
			So(suggestion.Action.(DialAction).PhoneNumber, ShouldEqual, "xyz")
		})

		Convey("View Location Action", func() {
			var suggestion Suggestion
			suggestion.SuggestedBase = SuggestedBase{
				Text:         "text",
				PostbackData: "postback data",
			}
			suggestion.Action = ViewLocationAction{
				LatLong: Location{
					Latitude:  4.12,
					Longitude: 2.23,
				},
				Label: "label",
				Query: "query",
			}

			b, err := json.Marshal(suggestion)
			So(err, ShouldBeNil)
			So(bytes.Compare(b, []byte(viewLocationActionJSON)), ShouldEqual, 0)

			err = json.Unmarshal(b, &suggestion)
			So(err, ShouldBeNil)
			So(suggestion.SuggestedBase.Text, ShouldEqual, "text")
			So(suggestion.SuggestedBase.PostbackData, ShouldEqual, "postback data")
			So(suggestion.Action, ShouldNotBeNil)
			So(suggestion.Action.(ViewLocationAction).Label, ShouldEqual, "label")
			So(suggestion.Action.(ViewLocationAction).Query, ShouldEqual, "query")
			So(suggestion.Action.(ViewLocationAction).LatLong, ShouldNotBeNil)
			So(suggestion.Action.(ViewLocationAction).LatLong.Latitude, ShouldEqual, 4.12)
			So(suggestion.Action.(ViewLocationAction).LatLong.Longitude, ShouldEqual, 2.23)
		})

		Convey("Create Calendar Event Action", func() {
			var suggestion Suggestion
			suggestion.SuggestedBase = SuggestedBase{
				Text:         "text",
				PostbackData: "postback data",
			}
			suggestion.Action = CreateCalendarEventAction{
				StartTime:   start,
				EndTime:     end,
				Title:       "title",
				Description: "description",
			}

			b, err := json.Marshal(suggestion)
			So(err, ShouldBeNil)
			So(bytes.Compare(b, []byte(createCalendarEventActionJSON)), ShouldEqual, 0)

			err = json.Unmarshal(b, &suggestion)
			So(err, ShouldBeNil)
			So(suggestion.SuggestedBase.Text, ShouldEqual, "text")
			So(suggestion.SuggestedBase.PostbackData, ShouldEqual, "postback data")
			So(suggestion.Action, ShouldNotBeNil)
			So(suggestion.Action.(CreateCalendarEventAction).Description, ShouldEqual, "description")
			So(suggestion.Action.(CreateCalendarEventAction).Title, ShouldEqual, "title")
			So(suggestion.Action.(CreateCalendarEventAction).StartTime, ShouldEqual, start)
			So(suggestion.Action.(CreateCalendarEventAction).EndTime, ShouldEqual, end)
		})

		Convey("Open URL Action", func() {
			var suggestion Suggestion
			suggestion.SuggestedBase = SuggestedBase{
				Text:         "text",
				PostbackData: "postback data",
			}
			var url = "https://www.google.com"
			suggestion.Action = OpenURLAction{
				URL: url,
			}

			b, err := json.Marshal(suggestion)
			So(err, ShouldBeNil)
			So(bytes.Compare(b, []byte(openURLActionJSON)), ShouldEqual, 0)

			err = json.Unmarshal(b, &suggestion)
			So(err, ShouldBeNil)
			So(suggestion.SuggestedBase.Text, ShouldEqual, "text")
			So(suggestion.SuggestedBase.PostbackData, ShouldEqual, "postback data")
			So(suggestion.Action, ShouldNotBeNil)
			So(suggestion.Action.(OpenURLAction).URL, ShouldEqual, url)
		})
	})
}
