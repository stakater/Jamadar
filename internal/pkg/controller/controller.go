package controller

import (
	"log"
	"time"

	"github.com/stakater/Jamadaar/internal/pkg/actions"
	"github.com/stakater/Jamadaar/internal/pkg/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

const jamadaarDisableAnnotation = "jamadaar.stakater.com/persist"

// Controller Jamadaar Controller to check for left over items
type Controller struct {
	clientset clientset.Interface
	config    config.Config
	Actions   []actions.Action
}

// NewController for initializing the Controller
func NewController(clientset clientset.Interface, config config.Config) (*Controller, error) {
	controller := &Controller{
		clientset: clientset,
		config:    config,
	}
	controller.Actions = actions.PopulateFromConfig(config.Actions)
	return controller, nil
}

//Run function for controller which handles the logic
func (c *Controller) Run() {
	for {
		nameSpaceList, err := c.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
		if err != nil {
			log.Printf("Error getting namespaces: %v", err)
			return
		}
		for _, namespace := range nameSpaceList.Items {
			annotations := namespace.Annotations
			if value, ok := annotations[jamadaarDisableAnnotation]; ok {
				if value != "true" {
					creationTime := namespace.CreationTimestamp
					weekAgoTime := time.Now().AddDate(0, 0, -7)
					weekAgo := metav1.NewTime(weekAgoTime)
					if creationTime.Before(&weekAgo) {
						c.clientset.CoreV1().Namespaces().Delete(namespace.Name, &metav1.DeleteOptions{})
						for _, action := range c.Actions {
							action.TakeAction(namespace)
						}
					}
				}
			}
		}
		timeInterval := c.config.TimeInterval
		duration, err := time.ParseDuration(timeInterval)
		time.Sleep(duration)
	}
}
