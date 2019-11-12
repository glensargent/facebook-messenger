package messenger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
		MessagingType: "UPDATE",
		Recipient:     recipient{ID: recipientID},
		Message:       textMessageContent{Text: text},
	}
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
	url := fmt.Sprintf("%v?access_token=%v", baseURL, c.AccessToken)
	msg, _ := json.Marshal(m)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return MsgResponse{}, err
	}

	return decode(resp)
}

// decodeResponse decodes Facebook response after sending message, usually contains MessageID or Error
func decode(r *http.Response) (MsgResponse, error) {
	defer r.Body.Close()
	var res MsgRawResponse
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return MsgResponse{}, err
	}

	// if the response has an error
	if res.Error != nil {
		return MsgResponse{}, errors.New(res.Error.Message)
	}

	return MsgResponse{
		MessageID:   res.MessageID,
		RecipientID: res.RecipientID,
	}, nil
}
