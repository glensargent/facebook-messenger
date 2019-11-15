package messenger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Persona represents a false user that talks to the recipient
// For more info, view here: https://developers.facebook.com/docs/messenger-platform/send-messages/personas
type Persona struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture_url"`
}

// GetPersonas retrieves all personas for this page
func (c Client) GetPersonas() ([]Persona, error) {
	url := fmt.Sprintf("https://graph.facebook.com/v5.0/me/personas/?access_token=%v", c.AccessToken)
	resp, err := httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var response PersonaResponse
	json.Unmarshal(body, &response)
	return response.Data, nil
}

// NewPersona creates a new Persona structure
func (c Client) NewPersona(name, pictureURL string) Persona {
	return Persona{
		Name:           name,
		ProfilePicture: pictureURL,
	}
}

// SavePersona stores the persona with the messenger API and returns that personas ID
func (c Client) SavePersona(pesrona Persona) string {
	return ""
}
