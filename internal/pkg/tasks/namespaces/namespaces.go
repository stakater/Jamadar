package namespaces

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/stakater/Jamadar/internal/pkg/actions"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

const jamadarDisableAnnotation = "jamadar.stakater.com/persist"

// DeleteNamespaces deletes the namespaces and takes the actions
func DeleteNamespaces(clientset clientset.Interface, namespaceList *v1.NamespaceList, actions []actions.Action, age string) error {

	for _, namespace := range namespaceList.Items {
		isDeleted, err := DeleteNamespace(clientset, namespace, age)
		if err != nil {
			return err
		}
		if isDeleted {
			log.Println("Namespace : " + namespace.Name + " deleted, Now Performing actions.")
			for _, action := range actions {
				action.TakeAction(namespace)
			}
		}
	}
	return nil
}

// DeleteNamespace deletes a single namespace
func DeleteNamespace(clientset clientset.Interface, namespace v1.Namespace, age string) (bool, error) {
	annotations := namespace.Annotations
	value, ok := annotations[jamadarDisableAnnotation]
	// check if annotation is not present and its value is not true
	if !ok || value != "true" {
		if checkIfOld(namespace, age) {
			err := clientset.CoreV1().Namespaces().Delete(namespace.Name, &metav1.DeleteOptions{})
			if err != nil {
				return false, err
			}
			return true, nil
		}
		return false, nil
	}
	return false, nil
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
