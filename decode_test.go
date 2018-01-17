package rcs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecode(t *testing.T) {
	Convey("Decode", t, func() {
		Convey("Decode Delivered Event Message", func() {
			req := []byte(`{
				"message": {
					"data": {
						"senderPhoneNumber": "+36702881817",
						"eventType": "DELIVERED",
						"eventId": "MsNLrzMfwjRF2uIwU9DBq0ZQ",
						"messageId": "a9b0e907-3580-4e61-a92e-2cb7d9530bc6",
						"sendTime": "2018-01-16T16:00:31.007983Z"
					},
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

			msg, err := Decode(req)
			So(err, ShouldBeNil)

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
						"data": {
							"senderPhoneNumber": "+36702881817",
							"eventType": "READ",
							"eventId": "MsvcQ0LKJRSLqe32sD-nfCfQ",
							"messageId": "a9b0e907-3580-4e61-a92e-2cb7d9530bc6",
							"sendTime": "2018-01-16T16:00:32.733778Z"
						},
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

			msg, err := Decode(req)
			So(err, ShouldBeNil)

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
						"data": {
							"senderPhoneNumber": "+36702881817",
							"eventType": "IS_TYPING",
							"eventId": "FKSCga1njmzehi_A",
							"sendTime": "2018-01-16T15:59:43.710048Z"
						},
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

			msg, err := Decode(req)
			So(err, ShouldBeNil)

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
						"data": {
							"senderPhoneNumber": "+36702881817",
							"messageId": "Ms4QPctbWeRgykVPjq9TH6vA",
							"sendTime": "2018-01-16T15:59:47.419510Z",
							"text": "😍"
						},
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

			msg, err := Decode(req)
			So(err, ShouldBeNil)

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
						"data": {
							"senderPhoneNumber": "+36702881817",
							"messageId": "MsWDRuvZvwSPurfMuPlWH2-Q",
							"sendTime": "2018-01-16T15:59:51.957968Z",
							"userFile": {
								"payload": {
									"mimeType": "image/gif",
									"fileSizeBytes": 162363,
									"fileName": "2130837663.gif",
									"fileUri": "https://storage.googleapis.com/gibe_rcs_copper_dev_test/c4d7d435-c24f-41b0-b736-e6f2eacf31d3/a47391cc57813de4bd93c3a0dc1bed8ea1ba1228455114b29931bfe52064"
								}
							}
						},
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

			msg, err := Decode(req)
			So(err, ShouldBeNil)

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
						"data": {
							"senderPhoneNumber": "+36702881817",
							"messageId": "Ms7TchISq9SCCRepoxMtHz4Q",
							"sendTime": "2018-01-16T15:59:54.845236Z",
							"location": {
								"latitude": 47.5025895,
								"longitude": 19.0492885
							}
						},
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

			msg, err := Decode(req)
			So(err, ShouldBeNil)

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
						"data": {
							"senderPhoneNumber": "+36702881817",
							"messageId": "Ms=Tf9X0YJQ1udXLQBegLAKg",
							"sendTime": "2018-01-16T16:00:34.689694Z",
							"suggestionResponse": {
								"postbackData": "alma",
								"text": "Hello"
							}
						},
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

			msg, err := Decode(req)
			So(err, ShouldBeNil)

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
