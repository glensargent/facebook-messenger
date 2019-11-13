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
	MessagingType string             `json:"messaging_type"`
	Recipient     recipient          `json:"recipient"`
	Message       textMessageContent `json:"message"`
}

type textMessageContent struct {
	Text string `json:"text"`
}

// NewTextMessage returns a new text message structure
func (c Client) NewTextMessage(recipientID int, text string) TextMessage {
	return TextMessage{
		MessagingType: "UPDATE",
		Recipient:     recipient{ID: recipientID},
		Message:       textMessageContent{Text: text},
	}
}
