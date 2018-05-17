package slack

import (
	"log"
	"testing"

	"github.com/stakater/Jamadar/internal/pkg/config"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	configFilePath   = "../../../../configs/testConfigs/CorrectSlackConfig.yaml"
	configuration, _ = config.ReadConfig(configFilePath)
)

type SlackMock struct {
}

func (s *SlackMock) SendNotification(message string) error {
	log.Print(message)
	return nil
}
func TestSlack_Init(t *testing.T) {
	type fields struct {
		Token   string
		Channel string
	}
	type args struct {
		params map[interface{}]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "MissingSlackToken",
			args: args{
				params: map[interface{}]interface{}{
					"token":   "",
					"channel": "channelName",
				},
			},
			wantErr: true,
		},
		{
			name: "CorrectScenario",
			args: args{
				params: map[interface{}]interface{}{
					"token":   "123",
					"channel": "channelName",
				},
			},
			fields: fields{
				Token:   "123",
				Channel: "channelName",
			},
		},

		{
			name: "ErrorInDecoding",
			args: args{
				params: map[interface{}]interface{}{
					"tokens":  "123",
					"channel": "channelName",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Slack{
				Token:   tt.fields.Token,
				Channel: tt.fields.Channel,
			}
			if err := s.Init(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Slack.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlack_TakeAction(t *testing.T) {
	type fields struct {
		Token   string
		Channel string
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TakeActionPass",
			args: args{
				obj: &v1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: "ns-test",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SlackMock{}
			message := "Namespace " + tt.args.obj.(*v1.Namespace).Name + " Deleted"
			s.SendNotification(message)
		})
	}
}
