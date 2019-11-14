package messenger

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestNewGenericMessage(t *testing.T) {
	msg := New("access_token", "123")

	expected := GenericMessage{
		Recipient: recipient{ID: 123},
		Message: genericMessageContent{
			Attachment: &attachment{
				Type:    "template",
				Payload: payload{TemplateType: "generic"},
			},
		},
	}

	got := msg.NewGenericMessage(123)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got %+v wanted %+v", got, expected)
	}
}

func TestAddElement(t *testing.T) {
	msg := New("access_token", "123")

	expected := GenericMessage{
		Recipient: recipient{ID: 123},
		Message: genericMessageContent{
			Attachment: &attachment{
				Type: "template",
				Payload: payload{
					TemplateType: "generic",
					Elements: []Element{
						Element{
							Title:    "Title",
							ImageURL: "image",
							Buttons: []Button{
								Button{
									Title: "btn title",
									Type:  "web_url",
									URL:   "test url",
								},
							},
						},
					},
				},
			},
		},
	}

	got := msg.NewGenericMessage(123)
	got.AddElement(Element{
		Title:    "Title",
		ImageURL: "image",
		Buttons: []Button{
			Button{
				Title: "btn title",
				Type:  "web_url",
				URL:   "test url",
			},
		},
	})

	if !reflect.DeepEqual(got.Message.Attachment.Payload.Elements, expected.Message.Attachment.Payload.Elements) {
		t.Errorf("got %+v wanted %+v", got, expected)
	}
}

func TestAddQuickReply(t *testing.T) {
	msg := New("access_token", "123")
	got := msg.NewGenericMessage(123)
	got.AddQuickReply(QuickReply{
		ContentType: "text",
		Title:       "Yes",
		Payload:     []byte("test"),
	})

	expected := GenericMessage{
		Recipient: recipient{ID: 123},
		Message: genericMessageContent{
			Attachment: &attachment{
				Type:    "template",
				Payload: payload{TemplateType: "generic"},
			},
			QuickReplies: []QuickReply{
				QuickReply{
					ContentType: "text",
					Title:       "Yes",
					Payload:     []byte("test"),
				},
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %+v wanted %+v", got, expected)
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
