package actions

import (
	"reflect"
	"testing"

	"github.com/stakater/Jamadaar/internal/pkg/config"
)

func TestPopulateFromConfig(t *testing.T) {
	type args struct {
		configActions []config.Action
	}
	tests := []struct {
		name string
		args args
		want []Action
	}{
		{
			name: "PopulateSlackAction",
			args: args{
				configActions: []config.Action{
					config.Action{
						Name: "slack",
						Params: map[interface{}]interface{}{
							"token":   "123",
							"channel": "channelName",
						},
					},
				},
			},
			want: []Action{
				MapToAction("slack"),
			},
		},
		{
			name: "PopulateSlackActionError",
			args: args{
				configActions: []config.Action{
					config.Action{
						Name: "slack",
						Params: map[interface{}]interface{}{
							"tok":  "123",
							"chan": "channelName",
						},
					},
				},
			},
			want: []Action{
				MapToAction("slack"),
			},
		},
		{
			name: "PopulateDefaultAction",
			args: args{
				configActions: []config.Action{
					config.Action{
						Name: "default",
					},
				},
			},
			want: []Action{
				MapToAction("default"),
			},
		},
		{
			name: "PopulateEmptyDefaultAction",
			args: args{
				configActions: []config.Action{},
			},
			want: []Action{
				MapToAction("default"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopulateFromConfig(tt.args.configActions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopulateFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
