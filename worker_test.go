package main

import (
	"errors"
	"testing"

	"myhttp/mocks"
)

func TestWorker_Run(t *testing.T) {
	type fields struct {
		number      int
		client      *mocks.ClientMock
		urlModifier *mocks.UrlModifierMock
	}
	type args struct {
		urls []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "No errors",
			fields: fields{
				number:      10,
				client:      mocks.NewClientMock(3, []byte{}, nil),
				urlModifier: mocks.NewUrlModifierMock(3, "http://url1", nil),
			},
			args: args{
				[]string{"url1", "url2", "url3"},
			},
		},
		{
			name: "Error on url modify but all requests performed",
			fields: fields{
				number:      10,
				client:      mocks.NewClientMock(3, []byte{}, nil),
				urlModifier: mocks.NewUrlModifierMock(3, "", errors.New("some error")),
			},
			args: args{
				[]string{"url1", "url2", "url3"},
			},
		},
		{
			name: "Error on client",
			fields: fields{
				number:      10,
				client:      mocks.NewClientMock(3, nil, errors.New("some error")),
				urlModifier: mocks.NewUrlModifierMock(3, "", errors.New("some error")),
			},
			args: args{
				[]string{"url1", "url2", "url3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := tt.fields.client
			modifier := tt.fields.urlModifier
			w := NewWorker(tt.fields.number, client, modifier)
			w.Run(tt.args.urls)
			if !client.AssertCallsCount() {
				t.Errorf("Client calls count assertion error: expected %d actual %d", client.ExpectedCallsCount(), client.CallsCount())
			}
			if !modifier.AssertCallsCount() {
				t.Errorf("Urls modifier calls count assertion error: expected %d actual %d", modifier.ExpectedCallsCount(), modifier.CallsCount())
			}
		})
	}
}
