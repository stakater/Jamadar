package controller

import (
	"github.com/stakater/Jamadaar/internal/pkg/config"
	clientset "k8s.io/client-go/kubernetes"
)

type Controller struct {
	clientset clientset.Interface
	config    config.Config
}

// NewController for initializing a Controller
func NewController(clientset clientset.Interface, config config.Config) (*Controller, error) {
	controller := &Controller{
		clientset: clientset,
		config:    config,
	}
	return controller, nil
}

//Run function for controller which handles the queue
func (c *Controller) Run() {

}
