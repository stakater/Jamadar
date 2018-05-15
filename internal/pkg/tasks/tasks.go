package tasks

import (
	"log"

	"github.com/stakater/Jamadaar/internal/pkg/actions"
	"github.com/stakater/Jamadaar/internal/pkg/tasks/namespaces"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

type Task struct {
	clientset clientset.Interface
	actions   []actions.Action
	age       string
}

func NewTask(clientSet clientset.Interface, actions []actions.Action, age string) *Task {
	return &Task{
		clientset: clientSet,
		actions:   actions,
		age:       age,
	}
}

// PerformTasks handles all the cleanup tasks
func (t *Task) PerformTasks() {
	t.performNamespaceDeletion()
}

// performNamespaceDeletion handles the deletion of namespaces
func (t *Task) performNamespaceDeletion() {
	log.Println("Starting to delete Namespaces")
	namespaceList, err := t.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Printf("Error getting namespaces: %v", err)
		return
	}

	err = namespaces.DeleteNamespaces(t.clientset, namespaceList, t.actions, t.age)
	if err != nil {
		log.Printf("Error deleting namespaces: %v", err)
		return
	}
}
