package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var token string = os.Getenv("TELEGRAM_TOKEN")
var url string = "https://api.telegram.org/bot" + token + "/sendMessage"

type Notifiable interface {
	Send() error
}

// Message represents a Telegram message.
type TelegramMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func NewTelegramMessage(message string) (*TelegramMessage, error) {
	client_id, err := strconv.ParseInt(os.Getenv("WEB_CHECKER_TELEGRAM_CLIENT_ID"), 10, 64)
	if err != nil {
		return nil, err
	}
	return &TelegramMessage{client_id, message}, nil
}

func (m *TelegramMessage) Send() error {
	payload, err := json.Marshal(*m)
	if err != nil {
		return err
	}
	response, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Println("could not close response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Send message failed: Status was %v", response.StatusCode)
	}
	return nil
}
