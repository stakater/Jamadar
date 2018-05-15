package slack

import (
	"errors"
	"log"

	"github.com/asiyani/slack"
	"github.com/mitchellh/mapstructure"
	"k8s.io/api/core/v1"
)

type SlackService interface {
	SendNotification(message string) error
}

// Slack action class implementing the Action interface
type Slack struct {
	Token   string
	Channel string
}

// Init initializes the Slack Configuration like token and channel
func (s *Slack) Init(params map[interface{}]interface{}) error {
	err := mapstructure.Decode(params, &s) //Converts the params to slack struct fields
	if err != nil {
		return err
	}
	if s.Token == "" || s.Channel == "" {
		return errors.New("Missing slack token or channel")
	}
	return nil
}

// TakeAction handles the main logic for slack action
func (s *Slack) TakeAction(obj interface{}) {
	message := "Namespace " + obj.(v1.Namespace).Name + " Deleted"
	err := s.SendNotification(message)
	if err != nil {
		log.Println("Error:  ", err)
	}
}

// SendNotification sends the Notification to the channel
func (s *Slack) SendNotification(message string) error {
	api := slack.New(s.Token)
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{prepareMessage(s, message)}
	params.AsUser = false

	_, _, err := api.PostMessage(s.Channel, "Jamadaar Alert", params)
	if err != nil {
		return err
	}

	log.Printf("Message successfully sent to Slack Channel `%s`", s.Channel)
	return nil
}

// Prepares the attachments to send in POST request
func prepareMessage(s *Slack, message string) slack.Attachment {
	return slack.Attachment{
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: message,
			},
		},
	}
}
