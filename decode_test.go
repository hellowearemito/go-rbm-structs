package rcs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	Sub = "projects/onyx-shoreline-191612/subscriptions/test-push-subscription"
)

func TestDecode(t *testing.T) {
	Convey("Decode", t, func() {
		Convey("Decode Delivered Event Message", func() {
			req := []byte(`{
				"message": {
					"data": "ew0KCQkJCQkJInNlbmRlclBob25lTnVtYmVyIjogIiszNjcwMjg4MTgxNyIsDQoJCQkJCQkiZXZlbnRUeXBlIjogIkRFTElWRVJFRCIsDQoJCQkJCQkiZXZlbnRJZCI6ICJNc05McnpNZndqUkYydUl3VTlEQnEwWlEiLA0KCQkJCQkJIm1lc3NhZ2VJZCI6ICJhOWIwZTkwNy0zNTgwLTRlNjEtYTkyZS0yY2I3ZDk1MzBiYzYiLA0KCQkJCQkJInNlbmRUaW1lIjogIjIwMTgtMDEtMTZUMTY6MDA6MzEuMDA3OTgzWiINCgkJCQkJfQ==",
					"attributes": {
						"event_type": "DELIVERED",
						"type": "event"
					},
					"message_id": "23352049337590",
					"messageId": "23352049337590",
					"publish_time": "2018-01-16T16:00:31.104Z",
					"publishTime": "2018-01-16T16:00:31.104Z"
				},
				"subscription": "projects/onyx-shoreline-191612/subscriptions/test-push-subscription"
			}`)

			subscription, msg, err := Decode(req)
			So(err, ShouldBeNil)
			So(subscription, ShouldNotBeNil)
			So(subscription, ShouldEqual, Sub)

			switch msg := msg.(type) {
			case DeliveredEventMessage:
				So(msg.Attributes.EventType, ShouldNotBeNil)
				So(*msg.Attributes.EventType, ShouldEqual, DeliveredEventType)
				So(msg.Attributes.Type, ShouldEqual, EventAttributeType)
				So(msg.MessageID, ShouldEqual, "23352049337590")
				So(msg.PublishTime.String(), ShouldEqual, "2018-01-16 16:00:31.104 +0000 UTC")
				So(msg.Data.SenderPhoneNumber, ShouldEqual, "+36702881817")
				So(msg.Data.EventType, ShouldEqual, DeliveredEventType)
				So(msg.Data.MessageID, ShouldEqual, "a9b0e907-3580-4e61-a92e-2cb7d9530bc6")
				So(msg.Data.SendTime.String(), ShouldEqual, "2018-01-16 16:00:31.007983 +0000 UTC")
			}
		})
		Convey("Decode Read event Message", func() {
			req := []byte(`
				{
					"message": {
						"data": "ew0KCQkJCQkJCSJzZW5kZXJQaG9uZU51bWJlciI6ICIrMzY3MDI4ODE4MTciLA0KCQkJCQkJCSJldmVudFR5cGUiOiAiUkVBRCIsDQoJCQkJCQkJImV2ZW50SWQiOiAiTXN2Y1EwTEtKUlNMcWUzMnNELW5mQ2ZRIiwNCgkJCQkJCQkibWVzc2FnZUlkIjogImE5YjBlOTA3LTM1ODAtNGU2MS1hOTJlLTJjYjdkOTUzMGJjNiIsDQoJCQkJCQkJInNlbmRUaW1lIjogIjIwMTgtMDEtMTZUMTY6MDA6MzIuNzMzNzc4WiINCgkJCQkJCX0=",
						"attributes": {
							"type": "event",
							"event_type": "READ"
						},
						"message_id": "23352128238996",
						"messageId": "23352128238996",
						"publish_time": "2018-01-16T16:00:32.915Z",
						"publishTime": "2018-01-16T16:00:32.915Z"
					},
					"subscription": "projects/onyx-shoreline-191612/subscriptions/test-push-subscription"
				}
			`)

			subscription, msg, err := Decode(req)
			So(err, ShouldBeNil)
			So(subscription, ShouldNotBeNil)
			So(subscription, ShouldEqual, Sub)

			switch msg := msg.(type) {
			case ReadEventMessage:
				So(msg.Attributes.EventType, ShouldNotBeNil)
				So(*msg.Attributes.EventType, ShouldEqual, ReadEventType)
				So(msg.Attributes.Type, ShouldEqual, EventAttributeType)
				So(msg.MessageID, ShouldEqual, "23352128238996")
				So(msg.PublishTime.String(), ShouldEqual, "2018-01-16 16:00:32.915 +0000 UTC")
				So(msg.Data.SenderPhoneNumber, ShouldEqual, "+36702881817")
				So(msg.Data.EventType, ShouldEqual, ReadEventType)
				So(msg.Data.MessageID, ShouldEqual, "a9b0e907-3580-4e61-a92e-2cb7d9530bc6")
				So(msg.Data.SendTime.String(), ShouldEqual, "2018-01-16 16:00:32.733778 +0000 UTC")
			}
		})
		Convey("Decode Is Typing Event Message", func() {
			req := []byte(`
				{
					"message": {
						"data": "ew0KCQkJCQkJCSJzZW5kZXJQaG9uZU51bWJlciI6ICIrMzY3MDI4ODE4MTciLA0KCQkJCQkJCSJldmVudFR5cGUiOiAiSVNfVFlQSU5HIiwNCgkJCQkJCQkiZXZlbnRJZCI6ICJGS1NDZ2ExbmptemVoaV9BIiwNCgkJCQkJCQkic2VuZFRpbWUiOiAiMjAxOC0wMS0xNlQxNTo1OTo0My43MTAwNDhaIg0KCQkJCQkJfQ==",
						"attributes": {
							"type": "event",
							"event_type": "IS_TYPING"
						},
						"message_id": "23352036536039",
						"messageId": "23352036536039",
						"publish_time": "2018-01-16T15:59:43.852Z",
						"publishTime": "2018-01-16T15:59:43.852Z"
					},
					"subscription": "projects/onyx-shoreline-191612/subscriptions/test-push-subscription"
				}
			`)

			subscription, msg, err := Decode(req)
			So(err, ShouldBeNil)
			So(subscription, ShouldNotBeNil)
			So(subscription, ShouldEqual, Sub)

			switch msg := msg.(type) {
			case IsTypingEventMessage:
				So(msg.Attributes.EventType, ShouldNotBeNil)
				So(*msg.Attributes.EventType, ShouldEqual, IsTypingEventType)
				So(msg.Attributes.Type, ShouldEqual, EventAttributeType)
				So(msg.MessageID, ShouldEqual, "23352036536039")
				So(msg.PublishTime.String(), ShouldEqual, "2018-01-16 15:59:43.852 +0000 UTC")
				So(msg.Data.SenderPhoneNumber, ShouldEqual, "+36702881817")
				So(msg.Data.EventType, ShouldEqual, IsTypingEventType)
				So(msg.Data.MessageID, ShouldBeEmpty)
				So(msg.Data.SendTime.String(), ShouldEqual, "2018-01-16 15:59:43.710048 +0000 UTC")
			}
		})
		Convey("Decode Text Message", func() {
			req := []byte(`
				{
					"message": {
						"data": "ew0KCQkJCQkJCSJzZW5kZXJQaG9uZU51bWJlciI6ICIrMzY3MDI4ODE4MTciLA0KCQkJCQkJCSJtZXNzYWdlSWQiOiAiTXM0UVBjdGJXZVJneWtWUGpxOVRINnZBIiwNCgkJCQkJCQkic2VuZFRpbWUiOiAiMjAxOC0wMS0xNlQxNTo1OTo0Ny40MTk1MTBaIiwNCgkJCQkJCQkidGV4dCI6ICLwn5iNIg0KCQkJCQkJfQ==",
						"attributes": {
							"type": "message",
							"message_type": "TEXT"
						},
						"message_id": "23351826002984",
						"messageId": "23351826002984",
						"publish_time": "2018-01-16T15:59:47.601Z",
						"publishTime": "2018-01-16T15:59:47.601Z"
					},
					"subscription": "projects/onyx-shoreline-191612/subscriptions/test-push-subscription"
				}	
			`)

			subscription, msg, err := Decode(req)
			So(err, ShouldBeNil)
			So(subscription, ShouldNotBeNil)
			So(subscription, ShouldEqual, Sub)

			switch msg := msg.(type) {
			case TextMessage:
				So(msg.Attributes.MessageType, ShouldNotBeNil)
				So(*msg.Attributes.MessageType, ShouldEqual, TextMessageType)
				So(msg.Attributes.Type, ShouldEqual, MessageAttributeType)
				So(msg.MessageID, ShouldEqual, "23351826002984")
				So(msg.PublishTime.String(), ShouldEqual, "2018-01-16 15:59:47.601 +0000 UTC")
				So(msg.Data.SenderPhoneNumber, ShouldEqual, "+36702881817")
				So(msg.Data.MessageID, ShouldEqual, "Ms4QPctbWeRgykVPjq9TH6vA")
				So(msg.Data.Text, ShouldEqual, "😍")
				So(msg.Data.SendTime.String(), ShouldEqual, "2018-01-16 15:59:47.41951 +0000 UTC")
			}
		})
		Convey("Decode User File Message", func() {
			req := []byte(`
				{
					"message": {
						"data": "ew0KCQkJCQkJCSJzZW5kZXJQaG9uZU51bWJlciI6ICIrMzY3MDI4ODE4MTciLA0KCQkJCQkJCSJtZXNzYWdlSWQiOiAiTXNXRFJ1dlp2d1NQdXJmTXVQbFdIMi1RIiwNCgkJCQkJCQkic2VuZFRpbWUiOiAiMjAxOC0wMS0xNlQxNTo1OTo1MS45NTc5NjhaIiwNCgkJCQkJCQkidXNlckZpbGUiOiB7DQoJCQkJCQkJCSJwYXlsb2FkIjogew0KCQkJCQkJCQkJIm1pbWVUeXBlIjogImltYWdlL2dpZiIsDQoJCQkJCQkJCQkiZmlsZVNpemVCeXRlcyI6IDE2MjM2MywNCgkJCQkJCQkJCSJmaWxlTmFtZSI6ICIyMTMwODM3NjYzLmdpZiIsDQoJCQkJCQkJCQkiZmlsZVVyaSI6ICJodHRwczovL3N0b3JhZ2UuZ29vZ2xlYXBpcy5jb20vZ2liZV9yY3NfY29wcGVyX2Rldl90ZXN0L2M0ZDdkNDM1LWMyNGYtNDFiMC1iNzM2LWU2ZjJlYWNmMzFkMy9hNDczOTFjYzU3ODEzZGU0YmQ5M2MzYTBkYzFiZWQ4ZWExYmExMjI4NDU1MTE0YjI5OTMxYmZlNTIwNjQiDQoJCQkJCQkJCX0NCgkJCQkJCQl9DQoJCQkJCQl9",
						"attributes": {
							"type": "message",
							"message_type": "USER_FILE"
						},
						"message_id": "23352089550964",
						"messageId": "23352089550964",
						"publish_time": "2018-01-16T15:59:52.121Z",
						"publishTime": "2018-01-16T15:59:52.121Z"
					},
					"subscription": "projects/onyx-shoreline-191612/subscriptions/test-push-subscription"
				}
			`)

			subscription, msg, err := Decode(req)
			So(err, ShouldBeNil)
			So(subscription, ShouldNotBeNil)
			So(subscription, ShouldEqual, Sub)

			switch msg := msg.(type) {
			case UserFileMessage:
				So(msg.Attributes.MessageType, ShouldNotBeNil)
				So(*msg.Attributes.MessageType, ShouldEqual, UserFileMessageType)
				So(msg.Attributes.Type, ShouldEqual, MessageAttributeType)
				So(msg.MessageID, ShouldEqual, "23352089550964")
				So(msg.PublishTime.String(), ShouldEqual, "2018-01-16 15:59:52.121 +0000 UTC")
				So(msg.Data.SenderPhoneNumber, ShouldEqual, "+36702881817")
				So(msg.Data.MessageID, ShouldEqual, "MsWDRuvZvwSPurfMuPlWH2-Q")
				So(msg.Data.UserFile, ShouldNotBeNil)
				So(msg.Data.UserFile.Payload, ShouldNotBeNil)
				So(msg.Data.UserFile.Payload.FileName, ShouldEqual, "2130837663.gif")
				So(msg.Data.UserFile.Payload.MimeType, ShouldEqual, "image/gif")
				So(msg.Data.UserFile.Payload.FileURI, ShouldEqual, "https://storage.googleapis.com/gibe_rcs_copper_dev_test/c4d7d435-c24f-41b0-b736-e6f2eacf31d3/a47391cc57813de4bd93c3a0dc1bed8ea1ba1228455114b29931bfe52064")
				So(msg.Data.UserFile.Payload.FileSizeBytes, ShouldEqual, 162363)
				So(msg.Data.SendTime.String(), ShouldEqual, "2018-01-16 15:59:51.957968 +0000 UTC")
			}
		})
		Convey("Decode Location Message", func() {
			req := []byte(`
				{
					"message": {
						"data": "ew0KCQkJCQkJCSJzZW5kZXJQaG9uZU51bWJlciI6ICIrMzY3MDI4ODE4MTciLA0KCQkJCQkJCSJtZXNzYWdlSWQiOiAiTXM3VGNoSVNxOVNDQ1JlcG94TXRIejRRIiwNCgkJCQkJCQkic2VuZFRpbWUiOiAiMjAxOC0wMS0xNlQxNTo1OTo1NC44NDUyMzZaIiwNCgkJCQkJCQkibG9jYXRpb24iOiB7DQoJCQkJCQkJCSJsYXRpdHVkZSI6IDQ3LjUwMjU4OTUsDQoJCQkJCQkJCSJsb25naXR1ZGUiOiAxOS4wNDkyODg1DQoJCQkJCQkJfQ0KCQkJCQkJfQ==",
						"attributes": {
							"type": "message",
							"message_type": "LOCATION"
						},
						"message_id": "23352052427276",
						"messageId": "23352052427276",
						"publish_time": "2018-01-16T15:59:54.970Z",
						"publishTime": "2018-01-16T15:59:54.970Z"
					},
					"subscription": "projects/onyx-shoreline-191612/subscriptions/test-push-subscription"
				}
			`)

			subscription, msg, err := Decode(req)
			So(err, ShouldBeNil)
			So(subscription, ShouldNotBeNil)
			So(subscription, ShouldEqual, Sub)

			switch msg := msg.(type) {
			case LocationMessage:
				So(msg.Attributes.MessageType, ShouldNotBeNil)
				So(*msg.Attributes.MessageType, ShouldEqual, LocationMessageType)
				So(msg.Attributes.Type, ShouldEqual, MessageAttributeType)
				So(msg.MessageID, ShouldEqual, "23352052427276")
				So(msg.PublishTime.String(), ShouldEqual, "2018-01-16 15:59:54.97 +0000 UTC")
				So(msg.Data.SenderPhoneNumber, ShouldEqual, "+36702881817")
				So(msg.Data.MessageID, ShouldEqual, "Ms7TchISq9SCCRepoxMtHz4Q")
				So(msg.Data.Location, ShouldNotBeNil)
				So(msg.Data.Location.Latitude, ShouldEqual, 47.5025895)
				So(msg.Data.Location.Longitude, ShouldEqual, 19.0492885)
				So(msg.Data.SendTime.String(), ShouldEqual, "2018-01-16 15:59:54.845236 +0000 UTC")
			}
		})

		Convey("Decode Suggestion Response Message", func() {
			req := []byte(`
				{
					"message": {
						"data": "ew0KCQkJCQkJCSJzZW5kZXJQaG9uZU51bWJlciI6ICIrMzY3MDI4ODE4MTciLA0KCQkJCQkJCSJtZXNzYWdlSWQiOiAiTXM9VGY5WDBZSlExdWRYTFFCZWdMQUtnIiwNCgkJCQkJCQkic2VuZFRpbWUiOiAiMjAxOC0wMS0xNlQxNjowMDozNC42ODk2OTRaIiwNCgkJCQkJCQkic3VnZ2VzdGlvblJlc3BvbnNlIjogew0KCQkJCQkJCQkicG9zdGJhY2tEYXRhIjogImFsbWEiLA0KCQkJCQkJCQkidGV4dCI6ICJIZWxsbyINCgkJCQkJCQl9DQoJCQkJCQl9",
						"attributes": {
							"message_type": "SUGGESTION_RESPONSE",
							"type": "message"
						},
						"message_id": "23351810877252",
						"messageId": "23351810877252",
						"publish_time": "2018-01-16T16:00:34.844Z",
						"publishTime": "2018-01-16T16:00:34.844Z"
					},
					"subscription": "projects/onyx-shoreline-191612/subscriptions/test-push-subscription"
				}	
			`)

			subscription, msg, err := Decode(req)
			So(err, ShouldBeNil)
			So(subscription, ShouldNotBeNil)
			So(subscription, ShouldEqual, Sub)

			switch msg := msg.(type) {
			case SuggestionResponseMessage:
				So(msg.Attributes.MessageType, ShouldNotBeNil)
				So(*msg.Attributes.MessageType, ShouldEqual, SuggestionResponseMessageType)
				So(msg.Attributes.Type, ShouldEqual, MessageAttributeType)
				So(msg.MessageID, ShouldEqual, "23351810877252")
				So(msg.PublishTime.String(), ShouldEqual, "2018-01-16 16:00:34.844 +0000 UTC")
				So(msg.Data.SenderPhoneNumber, ShouldEqual, "+36702881817")
				So(msg.Data.MessageID, ShouldEqual, "Ms=Tf9X0YJQ1udXLQBegLAKg")
				So(msg.Data.SuggestionResponse, ShouldNotBeNil)
				So(msg.Data.SuggestionResponse.PostbackData, ShouldEqual, "alma")
				So(msg.Data.SuggestionResponse.Text, ShouldEqual, "Hello")
				So(msg.Data.SendTime.String(), ShouldEqual, "2018-01-16 16:00:34.689694 +0000 UTC")
			}
		})
	})
}
