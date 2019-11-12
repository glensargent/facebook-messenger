package messenger

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
