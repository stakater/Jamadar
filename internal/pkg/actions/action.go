package actions

import (
	"log"

	"github.com/stakater/Jamadaar/internal/pkg/actions/slack"
	"k8s.io/api/core/v1"
)

func assertActionImplementations() {
	var _ Action = (*Default)(nil)
	var _ Action = (*slack.Slack)(nil)
}

// DefaultAction the name for default action name
const (
	DefaultAction = "default"
)

// Action interface so that other actions like slack can implement this
type Action interface {
	Init(map[interface{}]interface{}) error
	TakeAction(obj interface{})
}

// Default class with empty implementations for any action that we dont support currently
type Default struct {
}

// Init initializes handler configuration
// Do nothing for default handler
func (d *Default) Init(params map[interface{}]interface{}) error {
	return nil
}

// TakeAction the main business logic of Action
func (d *Default) TakeAction(obj interface{}) {
	message := "Namespace " + obj.(v1.Namespace).Name + " Deleted as it was a week old"
	log.Printf(message)
}
