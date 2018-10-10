package namespaces

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/stakater/Jamadar/internal/pkg/actions"
	"github.com/stakater/Jamadar/internal/pkg/config"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

const jamadarDisableAnnotation = "jamadar.stakater.com/persist"

// NamespaceToDelete represents the namespace object to be deleted
type NamespaceToDelete struct {
	clientset            clientset.Interface
	actions              []actions.Action
	age                  string
	restrictedNamespaces []string
}

// NewNamespaceToDelete creates a new NamespaceToDelete object
func NewNamespaceToDelete(clientSet clientset.Interface, actions []actions.Action, conf config.Config) *NamespaceToDelete {
	return &NamespaceToDelete{
		clientset:            clientSet,
		actions:              actions,
		age:                  conf.Age,
		restrictedNamespaces: conf.RestrictedNamespaces,
	}
}

// DeleteNamespaces deletes the namespaces and takes the actions
func (n *NamespaceToDelete) DeleteNamespaces(namespaceList *v1.NamespaceList) error {

	for _, namespace := range namespaceList.Items {
		isDeleted, err := n.DeleteNamespace(namespace)
		if err != nil {
			return err
		}
		if isDeleted {
			log.Println("Namespace : " + namespace.Name + " deleted, Now Performing actions.")
			for _, action := range n.actions {
				action.TakeAction(namespace)
			}
		}
	}
	return nil
}

// DeleteNamespace deletes a single namespace
func (n *NamespaceToDelete) DeleteNamespace(namespace v1.Namespace) (bool, error) {
	annotations := namespace.Annotations
	value, ok := annotations[jamadarDisableAnnotation]
	// check if annotation is not present and its value is not true
	if !ok || value != "true" {
		if !n.isNamespaceinRestrictedNamespaces(namespace.Name) {
			if checkIfOld(namespace, n.age) {
				err := n.clientset.CoreV1().Namespaces().Delete(namespace.Name, &metav1.DeleteOptions{})
				if err != nil {
					return false, err
				}
				return true, nil
			}
		}
		return false, nil
	}
	return false, nil
}

func (n *NamespaceToDelete) isNamespaceinRestrictedNamespaces(namespace string) bool {
	for _, restrictedNamespace := range n.restrictedNamespaces {
		if restrictedNamespace == namespace {
			return true
		}
	}
	return false

}

// checkIfOld checks if the namespace is 7 days old
func checkIfOld(namespace v1.Namespace, age string) bool {
	creationTime := namespace.CreationTimestamp
	var day, month, year int
	if age[len(age)-1] == 'd' {
		age := strings.TrimSuffix(age, "d")
		days, _ := strconv.Atoi(age)
		day = -1 * days
		month = 0
		year = 0
	} else if age[len(age)-1] == 'w' {
		age := strings.TrimSuffix(age, "w")
		weeks, _ := strconv.Atoi(age)
		day = -1 * 7 * weeks
		month = 0
		year = 0
	} else if age[len(age)-1] == 'm' {
		age := strings.TrimSuffix(age, "m")
		months, _ := strconv.Atoi(age)
		day = 0
		month = -1 * months
		year = 0
	} else if age[len(age)-1] == 'y' {
		age := strings.TrimSuffix(age, "y")
		years, _ := strconv.Atoi(age)
		day = 0
		month = 0
		year = -1 * years
	}

	weekAgoTime := time.Now().AddDate(year, month, day)
	weekAgo := metav1.NewTime(weekAgoTime)
	return creationTime.Before(&weekAgo)
}
