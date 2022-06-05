package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestNotify(t *testing.T) {
	message := "this is the notification message"
	notifier := &bytes.Buffer{}

	err := Notify(message, notifier)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	got := notifier.String()
	if message != got {
		t.Fatalf("wanted %q, got %q", message, got)
	}
}

const client_id_mock = 123456789
const token_mock = "123456789:ABCDEFGHILMNOXYZ"
const message_mock = "telegram bot message test"

func TestNewTelegramNotify(t *testing.T) {
	os.Setenv("TELEGRAM_TOKEN", token_mock)
	os.Setenv("WEB_CHECKER_TELEGRAM_CLIENT_ID", fmt.Sprintf("%d", client_id_mock))
	got, err := NewTelegramNotify()
	want := TelegramNotify{telegram_api, client_id_mock, token_mock}

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

}

func TestSendTelegramMockOK(t *testing.T) {
	ts := NewTestServer(t, http.StatusOK)
	defer ts.Close()

	notifier := &TelegramNotify{ts.URL, client_id_mock, token_mock}

	err := Notify(message_mock, notifier)
	if err != nil {
		t.Fatalf("unexpected error %q", err)
	}
}

func TestSendTelegramMockError(t *testing.T) {
	ts := NewTestServer(t, http.StatusBadRequest)
	defer ts.Close()

	notifier := &TelegramNotify{ts.URL, client_id_mock, token_mock}
	err := Notify(message_mock, notifier)
	if err == nil {
		t.Fatalf("expected error got nil")
	}
}

func GetTelegramMessageFromRequest(t *testing.T, body io.ReadCloser) TelegramMessage {
	t.Helper()

	defer body.Close()
	out, err := ioutil.ReadAll(body)
	if err != nil {
		t.Errorf("could not read request's body: %q", err)
	}
	msg := TelegramMessage{}
	json.Unmarshal(out, &msg)
	return msg
}

func NewTestServer(t *testing.T, statusCode int) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		if statusCode != http.StatusOK {
			// not going to test the request content, just handling the returned error code
			return
		}

		if r.Method != "POST" {
			t.Errorf("Expected 'POST' request, was '%s'", r.Method)
		}

		want := fmt.Sprintf("/bot%s/sendMessage", token_mock)
		if r.URL.EscapedPath() != want {
			t.Errorf("url check: want '%s', got '%s'", want, r.URL.EscapedPath())
		}

		msg := GetTelegramMessageFromRequest(t, r.Body)
		if msg.ChatID != client_id_mock {
			t.Errorf("chat id check: expected client id %d, got %d", client_id_mock, msg.ChatID)
		}
		if msg.Text != message_mock {
			t.Errorf("text field: want '%s', got '%s'", message_mock, msg.Text)
		}
	}))
}
