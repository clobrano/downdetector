package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Notify(message string, writer io.Writer) error {
	_, err := fmt.Fprint(writer, message)
	return err
}

const telegram_api = "https://api.telegram.org"

type TelegramNotify struct {
	baseurl string
	chatId  int64
	token   string
}

func NewTelegramNotify() (TelegramNotify, error) {
	client_id, err := strconv.ParseInt(os.Getenv("WEB_CHECKER_TELEGRAM_CLIENT_ID"), 10, 64)
	if err != nil {
		return TelegramNotify{}, err
	}
	token := os.Getenv("TELEGRAM_TOKEN")
	return TelegramNotify{telegram_api, client_id, token}, nil
}

// Message represents a Telegram message.
type TelegramMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func (t *TelegramNotify) Write(p []byte) (n int, err error) {
	payload, err := json.Marshal(TelegramMessage{t.chatId, string(p)})
	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf("%s/bot%s/sendMessage", t.baseurl, t.token)
	response, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Send message failed: Status was %v (%s)", response.StatusCode, url)
	}
	return len(p), nil
}
