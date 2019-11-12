package messenger

import (
	"encoding/json"
	"errors"
	"net/http"
)

// MsgResponse received from Messenger API after sending the message
type MsgResponse struct {
	MessageID   string `json:"message_id"`
	RecipientID int64  `json:"recipient_id,string"`
}

// MsgRawResponse received from Messenger API after sending the message
// this caters to both success and errors, we use this to determine what type
// of response to give back to the user, on success, the error is removed and this becomes MsgResponse
type MsgRawResponse struct {
	MessageID   string    `json:"message_id"`
	RecipientID int64     `json:"recipient_id,string"`
	Error       *MsgError `json:"error"`
}

// MsgError received from Messenger API if sending has failed
type MsgError struct {
	Code      int    `json:"code"`
	FbtraceID string `json:"fbtrace_id"`
	Message   string `json:"message"`
	Type      string `json:"type"`
}

// decodes Messenger response after sending message and returns the correct
// structure based on the response having an error object or not
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
