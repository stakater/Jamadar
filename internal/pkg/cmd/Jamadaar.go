package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/stakater/Jamadaar/internal/pkg/config"
	"github.com/stakater/Jamadaar/internal/pkg/controller"
	"github.com/stakater/Jamadaar/pkg/kube"
)

func NewJamadaarCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "jamadaar",
		Short: "A watcher for your Kubernetes cluster",
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
	config := getControllerConfig()

	controller, err := controller.NewController(clientset, config)
	if err != nil {
		log.Printf("Error occured while creating controller. Reason: %s", err.Error())
	}
	// Now let's start the controller
	go controller.Run()

	// Wait forever
	select {}

}

// get the yaml configuration for the controller
func getControllerConfig() config.Config {
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if len(configFilePath) == 0 {
		//Default config file is placed in configs/ folder
		configFilePath = "configs/config.yaml"
	}
	configuration, err := config.ReadConfig(configFilePath)
	if err != nil {
		log.Panic(err)
	}
	return configuration
}
