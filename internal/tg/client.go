package tg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TgClient struct {
	token   string
	baseUrl string
}

func NewTgClient(token string) *TgClient {
	baseUrl := fmt.Sprintf("https://api.telegram.org/bot%s/", token)

	return &TgClient{token, baseUrl}
}

type sendMessageRequestBody struct {
	Text   string `json:"text"`
	ChatId string `json:"chat_id"`
}

func (tc *TgClient) SendMessage(text string, chatId string) error {
	body := sendMessageRequestBody{
		Text:   text,
		ChatId: chatId,
	}
	bodyJson, _ := json.Marshal(body)

	url := tc.baseUrl + "sendMessage"
	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("status code is not 200, got %d", res.StatusCode)
}
