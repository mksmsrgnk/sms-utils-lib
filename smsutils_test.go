package smsutils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTextMessage_send(t *testing.T) {
	type args struct {
		userName string
		password string
		rawURL   string
		from     string
		to       string
		text     string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1",
			args: args{
				userName: "test",
				password: "test",
				rawURL:   "http://localhost.localdomain",
				from:     "1040",
				to:       "+996555555555",
				text:     "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(
				http.HandlerFunc(
					func(
						w http.ResponseWriter,
						r *http.Request) {
						w.WriteHeader(202)
					}))
			defer ts.Close()
			err := NewKannel(
				tt.args.userName,
				tt.args.password,
				ts.URL).NewTextMessage(
				tt.args.from,
				tt.args.to,
				tt.args.text).Send()
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestTextMessage_encodeURL(t1 *testing.T) {
	type fields struct {
		Kannel Kannel
		From   string
		To     string
		Text   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{name: "test", fields: fields{
			Kannel: Kannel{
				URL:      "http://localhost.localdomain",
				UserName: "test",
				Password: "test"},
			From: "1040",
			To:   "+996555555555",
			Text: "test text",
		}, want: "http://localhost.localdomain?from=1040&password=test&text=test+text&to=%2B996555555555&username=test", wantErr: false},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TextMessage{
				Kannel: tt.fields.Kannel,
				From:   tt.fields.From,
				To:     tt.fields.To,
				Text:   tt.fields.Text,
			}
			got, err := t.encodeURL()
			if (err != nil) != tt.wantErr {
				t1.Errorf("encodeURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("encodeURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
