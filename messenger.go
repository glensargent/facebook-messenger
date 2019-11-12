package messenger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// BaseURL used for messenger API
var BaseURL = "https://graph.facebook.com/v5.0/me/messages/"

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

// SendMessage takes any type of message and posts to messenger API
func (c Client) SendMessage(m Message) (MsgResponse, error) {
	// type switch to check that the Message is a supported type
	switch m.(type) {
	case TextMessage:
	default:
		log.Println("Unsupported message type")
		return MsgResponse{}, nil
	}

	// construct the post request
	url := fmt.Sprintf("%v?access_token=%v", BaseURL, c.AccessToken)
	msg, _ := json.Marshal(m)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return MsgResponse{}, err
	}

	return decode(resp)
}

// SendTextMessage is a wrapper method that creates a text message type
// and sends that message for you
func (c Client) SendTextMessage(recipient int, msg string) {
	m := c.NewTextMessage(recipient, msg)
	c.SendMessage(m)
}
