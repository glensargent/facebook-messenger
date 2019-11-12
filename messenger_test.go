package messenger

import "testing"

func TestNewClient(t *testing.T) {
	// client := messenger.New("")
	got := New("123", "456")
	want := Client{"123", "456"}
	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func TestNewTextMessage(t *testing.T) {
	// client := messenger.New("")
	client := New("123", "456")
	got := client.NewTextMessage(123, "test")
	want := TextMessage{
		Recipient: recipient{ID: 123},
		Message:   textMessageContent{Text: "test"},
	}

	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}
