package smsutils

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	Address  string
	UserName string
	Password string
)

func format(text string) string {
	return strings.Replace(text, " ", "+", -1)
}

func prepareRequest(from, to, text string) string {
	return fmt.Sprintf("http://%s/cgi-bin/sendsms?username=%s&password=%s&from=%s&to=%s&text=%s",
		Address, UserName, Password, from, to, format(text))
}

func Send(from, to, text string) error {
	for _, mobPhone := range strings.Split(to, ",") {
		resp, err := http.Get(prepareRequest(from, to, text))
		if err != nil {
			return fmt.Errorf("SMS message sending error, %v", err)
		}
		if resp.StatusCode != http.StatusAccepted {
			return fmt.Errorf("can't send message to: %s, response code %d", mobPhone, resp.StatusCode)
		}
	}
	return nil
}
