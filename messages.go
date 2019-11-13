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

// GenericMessage struct used for sending structural messages to messenger (messages with images, links, and buttons)
type GenericMessage struct {
	Recipient recipient             `json:"recipient"`
	Message   genericMessageContent `json:"message"`
}
type genericMessageContent struct {
	Attachment *attachment `json:"attachment,omitempty"`
}

type attachment struct {
	Type    string  `json:"type,omitempty"`
	Payload payload `json:"payload,omitempty"`
}
type payload struct {
	TemplateType string    `json:"template_type,omitempty"`
	Elements     []Element `json:"elements,omitempty"`
}

// Element in Generic Message template
// https://developers.facebook.com/docs/messenger-platform/send-messages/template/generic
type Element struct {
	Title    string   `json:"title"`
	Subtitle string   `json:"subtitle,omitempty"`
	ItemURL  string   `json:"item_url,omitempty"`
	ImageURL string   `json:"image_url,omitempty"`
	Buttons  []Button `json:"buttons,omitempty"`
}

// Button on Generic Message template element
// https://developers.facebook.com/docs/messenger-platform/send-messages/template/generic
type Button struct {
	Type    string `json:"type"`
	URL     string `json:"url,omitempty"`
	Title   string `json:"title"`
	Payload string `json:"payload,omitempty"`
}

// NewTextMessage returns a new text message structure
func (c Client) NewTextMessage(recipientID int, text string) TextMessage {
	return TextMessage{
		MessagingType: "UPDATE",
		Recipient:     recipient{ID: recipientID},
		Message:       textMessageContent{Text: text},
	}
}

// NewGenericMessage creates new Generic Template message that's used for attaching other elements such as images, links, buttons etc
// https://developers.facebook.com/docs/messenger-platform/send-messages/template/generic
func (c Client) NewGenericMessage(recipientID int) GenericMessage {
	return GenericMessage{
		Recipient: recipient{ID: recipientID},
		Message: genericMessageContent{
			Attachment: &attachment{
				Type:    "template",
				Payload: payload{TemplateType: "generic"},
			},
		},
	}
}

// AddElement adds a new element to the message object
// https://developers.facebook.com/docs/messenger-platform/send-messages/template/generic
func (m *GenericMessage) AddElement(e Element) {
	m.Message.Attachment.Payload.Elements = append(m.Message.Attachment.Payload.Elements, e)
}
