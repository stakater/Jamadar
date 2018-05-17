package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/stakater/Jamadar/internal/pkg/config"
	"github.com/stakater/Jamadar/internal/pkg/controller"
	"github.com/stakater/Jamadar/pkg/kube"
)

//NewJamadarCommand to start and run Jamadar
func NewJamadarCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "jamadar",
		Short: "A kubernetes controller which cleans up left overs",
		Run:   startJamadar,
	}
	return cmds
}

func startJamadar(cmd *cobra.Command, args []string) {
	log.Println("Starting Jamadar")
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
