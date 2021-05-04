package main

import (
	"testing"

	gonsq "github.com/nsqio/go-nsq"
	"github.com/reyhanfahlevi/pkg/go/nsq"
)

func Test_handlerSample(t *testing.T) {
	type args struct {
		message nsq.IMessage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				message: &nsq.Message{
					Message: &gonsq.Message{
						Body:     []byte("{}"),
						Attempts: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handlerSample(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("handlerSample() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
