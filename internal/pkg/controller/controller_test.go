package controller

import (
	"testing"

	"github.com/stakater/Jamadar/internal/pkg/config"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestControllerPass(t *testing.T) {
	configuration := config.Config{
		PollTimeInterval: "1s",
		Age:              "7d",
		Actions: []config.Action{
			config.Action{
				Name: "default",
			},
		},
	}
	clientset := testclient.NewSimpleClientset()
	namespace := v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ns-test",
			Annotations: map[string]string{
				"jamadar.stakater.com/persist": "false",
			},
		},
	}
	clientset.CoreV1().Namespaces().Create(&namespace)
	controller, _ := NewController(clientset, configuration)
	controller.handleTasks()
}
