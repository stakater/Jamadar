package controller

import (
	"log"
	"time"

	"github.com/stakater/Jamadar/internal/pkg/actions"
	"github.com/stakater/Jamadar/internal/pkg/config"
	"github.com/stakater/Jamadar/internal/pkg/tasks"
	clientset "k8s.io/client-go/kubernetes"
)

// Controller Jamadar Controller to check for left over items
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
		c.handleTasks()
		timeInterval := c.config.PollTimeInterval
		duration, err := time.ParseDuration(timeInterval)
		if err != nil {
			log.Printf("Error Parsing Time Interval: %v", err)
			return
		}
		time.Sleep(duration)
	}
}

func (c *Controller) handleTasks() {
	task := tasks.NewTask(c.clientset, c.Actions, c.config)
	task.PerformTasks()
}
