package messenger

// Message represents any form of message structure to post to Facebook
// https://developers.facebook.com/docs/messenger-platform/reference/send-api/#message
type Message interface{}

// Recipient represents a user to send a message to
type recipient struct {
	ID int `json:"id"` // the recipients ID
}

// TextMessage is the structure that represents a text only message on messenger
type TextMessage struct {
	Recipient recipient `json:"recipient"`
	Message   textMessageContent
}

type textMessageContent struct {
	Text string `json:"text"`
}