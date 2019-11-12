package messenger

import (
	"net/http"
	"time"
)

// baseURL used for messenger API
const baseURL = "https://graph.facebook.com/v5.0/me/messages/"

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

// Client structure used to send and receive messages
type Client struct {
	AccessToken string
	PageID      string
}

// New returns a new messenger client
func New(accessToken, pageID string) Client {
	return Client{accessToken, pageID}
}

// NewTextMessage returns a new text message structure
func (c Client) NewTextMessage(recipientID int, text string) TextMessage {
	return TextMessage{
		Recipient: recipient{ID: recipientID},
		Message:   textMessageContent{Text: text},
	}
}

// SendMessage takes any type of message and posts to messenger API
func (c Client) SendMessage(m Message) {

}
