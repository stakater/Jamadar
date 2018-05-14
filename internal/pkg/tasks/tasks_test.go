package tasks

import (
	"testing"

	"github.com/stakater/Jamadaar/internal/pkg/actions"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestPerformTasks(t *testing.T) {
	type args struct {
		clientset clientset.Interface
		actions   []actions.Action
		age       string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "PerformTasksPass",
			args: args{
				clientset: testclient.NewSimpleClientset(),
				actions: []actions.Action{
					&actions.Default{},
				},
				age: "1d",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			namespace := v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "ns-test",
					Annotations: map[string]string{
						"jamadaar.stakater.com/persist": "false",
					},
				},
			}
			tt.args.clientset.CoreV1().Namespaces().Create(&namespace)
			PerformTasks(tt.args.clientset, tt.args.actions, tt.args.age)
		})
	}
}

func TestPerformTasksNoNamespaces(t *testing.T) {
	actions := []actions.Action{
		&actions.Default{},
	}
	PerformTasks(testclient.NewSimpleClientset(), actions, "1y")
}
