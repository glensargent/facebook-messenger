package messenger

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	got := New("123", "456")
	want := Client{"123", "456"}
	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func TestNewTextMessage(t *testing.T) {
	client := New("123", "456")
	got := client.NewTextMessage(123, "test")
	want := TextMessage{
		MessagingType: "UPDATE",
		Recipient:     recipient{ID: 123},
		Message:       textMessageContent{Text: "test"},
	}

	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func TestSendMessage(t *testing.T) {
	// fs will mock up fb messenger server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := MsgResponse{
			RecipientID: 123,
			MessageID:   "TEST123",
		}
		b, _ := json.Marshal(rec)
		w.Write(b)
	}))
	defer server.Close()

	BaseURL = server.URL

	client := New("123", "456")
	m := client.NewTextMessage(123, "test")
	got, err := client.SendMessage(m)

	if err != nil {
		t.Error(err)
	}

	expected := MsgResponse{
		MessageID:   "TEST123",
		RecipientID: 123,
	}

	if got != expected {
		t.Errorf("got %+v wanted %+v", got, expected)
	}
}
