package smsutils

import "testing"

func Test_sendSMS(t *testing.T) {
	type args struct {
		mobileNum string
		text      string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{mobileNum: "+996555555555", text: "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_convertMessageText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Message format conversion test", args{text: "message body"}, "message+body"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := format(tt.args.text); got != tt.want {
				t.Errorf("format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prepareRequest(t *testing.T) {
	type args struct {
		from string
		to   string
		text string
	}

	UserName = "test"
	Password = "test"
	Address = "127.0.0.1"

	tests := []struct {
		name string
		args args
		want string
	}{
		{"Kannel request string", args{from: "1040", to: "+996555555555", text: "test message"}, "http://127.0.0.1/cgi-bin/sendsms?username=test&password=test&from=1040&to=+996555555555&text=test+message",},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareRequest(tt.args.from, tt.args.to, tt.args.text); got != tt.want {
				t.Errorf("prepareRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
