package smsutils

import (
	"fmt"
	"net/http"
	"net/url"
)

type Kannel struct {
	URL      string
	UserName string
	Password string
}

type TextMessage struct {
	Kannel
	From string
	To   string
	Text string
}

func NewKannel(userName, password, rawURL string) Kannel {
	return Kannel{UserName: userName, Password: password, URL: rawURL}
}

func (k Kannel) NewTextMessage(from, to, text string) TextMessage {
	return TextMessage{From: from, To: to, Text: text, Kannel: k}
}

func (t TextMessage) encodeURL() (string, error) {
	u, err := url.Parse(t.URL)
	if err != nil {
		return "", err
	}
	params := url.Values{
		"username": []string{t.UserName},
		"password": []string{t.Password},
		"from":     []string{t.From},
		"to":       []string{t.To},
		"text":     []string{t.Text},
	}
	u.RawQuery = params.Encode()
	return u.String(), nil
}

func (t TextMessage) send(rawURL string) error {
	resp, err := http.Get(rawURL)
	if err != nil {
		return fmt.Errorf("can't send message: %v", err)
	}
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("can't send message to: %s, response code %d", t.To, resp.StatusCode)
	}
	return nil
}

func (t TextMessage) Send() error {
	req, err := t.encodeURL()
	if err != nil {
		return err
	}
	return t.send(req)
}
