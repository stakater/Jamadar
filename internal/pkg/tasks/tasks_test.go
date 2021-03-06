package tasks

import (
	"testing"

	"github.com/stakater/Jamadar/internal/pkg/actions"
	"github.com/stakater/Jamadar/internal/pkg/config"
	clientset "k8s.io/client-go/kubernetes"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestTask_PerformTasks(t *testing.T) {
	type fields struct {
		clientset clientset.Interface
		actions   []actions.Action
		conf      config.Config
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "PerformTasksPass",
			fields: fields{
				clientset: testclient.NewSimpleClientset(),
				actions: []actions.Action{
					&actions.Default{},
				},
				conf: config.Config{
					Age: "1d",
					Resources: []string{
						"namespaces",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &Task{
				clientset: tt.fields.clientset,
				actions:   tt.fields.actions,
				config:    tt.fields.conf,
			}
			task.PerformTasks()
		})
	}
}
func TestPerformTasksNoNamespaces(t *testing.T) {
	actions := []actions.Action{
		&actions.Default{},
	}
	conf := config.Config{
		Age: "1y",
		Resources: []string{
			"namespaces",
		},
	}
	task := NewTask(testclient.NewSimpleClientset(), actions, conf)
	task.PerformTasks()
}

func TestPerformDefault(t *testing.T) {
	actions := []actions.Action{
		&actions.Default{},
	}
	conf := config.Config{
		Age: "1y",
		Resources: []string{
			"default",
		},
	}
	task := NewTask(testclient.NewSimpleClientset(), actions, conf)
	task.PerformTasks()
}
