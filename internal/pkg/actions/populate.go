package actions

import (
	"log"

	"github.com/stakater/Jamadaar/internal/pkg/actions/slack"
	"github.com/stakater/Jamadaar/internal/pkg/config"
)

// PopulateFromConfig populates the actions for a specific controller from config
func PopulateFromConfig(configActions []config.Action) []Action {
	var populatedActions []Action
	if len(configActions) == 0 {
		configActions = []config.Action{
			config.Action{
				Name: "default",
			},
		}
	}
	for _, configAction := range configActions {
		actionToAdd := MapToAction(configAction.Name)
		err := actionToAdd.Init(configAction.Params)
		if err != nil {
			log.Println(err)
		}
		populatedActions = append(populatedActions, actionToAdd)
	}
	return populatedActions
}

// MapToAction maps the action name to the actual action type
func MapToAction(actionName string) Action {
	action, ok := actionMap[actionName]
	if !ok {
		return actionMap[DefaultAction]
	}
	return action
}

var actionMap = map[string]Action{
	"default": &Default{},
	"slack":   &slack.Slack{},
}
