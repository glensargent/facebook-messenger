package messenger

const (
	// TypingOn used for SenderAction
	TypingOn = "typing_on"
	// TypingOff used for SenderAction
	TypingOff = "typing_off"
	// MarkSeen used for SenderAction
	MarkSeen = "mark_seen"
)

// SenderAction used primarily for setting typing on and off
// for more info: https://developers.facebook.com/docs/messenger-platform/send-messages/sender-actions
type SenderAction struct {
	Recipient recipient `json:"recipient"`
	Action    string    `json:"sender_action"`
	PersonaID string    `json:"persona_id,omitempty"`
}

// NewAction creates a new user action
func (c Client) NewAction(recipientID int, action string) SenderAction {
	return SenderAction{
		Recipient: recipient{ID: recipientID},
		Action:    action,
	}
}

// SetTypingOn for the bot in messenger
func (c Client) SetTypingOn(recipientID int) (MsgResponse, error) {
	action := c.NewAction(recipientID, TypingOn)
	res, err := c.SendAction(action)
	if err != nil {
		return MsgResponse{}, err
	}

	return res, nil
}

// SetTypingOff for the bot in messenger
func (c Client) SetTypingOff(recipientID int) (MsgResponse, error) {
	action := c.NewAction(recipientID, TypingOff)
	res, err := c.SendAction(action)
	if err != nil {
		return MsgResponse{}, err
	}

	return res, nil
}

// MarkSeen for the bot in messenger
func (c Client) MarkSeen(recipientID int) (MsgResponse, error) {
	action := c.NewAction(recipientID, MarkSeen)
	res, err := c.SendAction(action)
	if err != nil {
		return MsgResponse{}, err
	}

	return res, nil
}

// AddPersona adds a persona ID to an action
func (s *SenderAction) AddPersona(personaID string) {
	s.PersonaID = personaID
}
