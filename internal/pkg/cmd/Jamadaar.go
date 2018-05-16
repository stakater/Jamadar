package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/stakater/Jamadaar/internal/pkg/config"
	"github.com/stakater/Jamadaar/internal/pkg/controller"
	"github.com/stakater/Jamadaar/pkg/kube"
)

//NewJamadaarCommand to start and run Jamadaar
func NewJamadaarCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "jamadaar",
		Short: "A kubernetes controller which cleans up left overs",
		Run:   startJamadaar,
	}
	return cmds
}

func startJamadaar(cmd *cobra.Command, args []string) {
	log.Println("Starting Jamadaar")
	// create the clientset
	clientset, err := kube.GetClient()
	if err != nil {
		log.Fatal(err)
	}

	// get the Controller config file
	config := config.GetConfiguration()

	controller, err := controller.NewController(clientset, config)
	if err != nil {
		log.Printf("Error occured while creating controller. Reason: %s", err.Error())
	}

	go controller.Run()

	// Wait forever
	select {}
}
