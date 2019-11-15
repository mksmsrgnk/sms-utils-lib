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
		if resp, err := http.Get(prepareRequest(from, to, text)); err == nil {
			if resp.StatusCode != http.StatusAccepted {
				return fmt.Errorf("can't send message to: %s, response code %d", mobPhone, resp.StatusCode)
			}
			return fmt.Errorf("message to: %s has been accepted for delivery!, %d", mobPhone, resp.StatusCode)
		} else {
			return fmt.Errorf("SMS message sending error, %v", err)
		}
	}
	return nil
}
