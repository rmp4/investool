package main

import (
	"fmt"
	"investool/pkg/config"
	"investool/pkg/log"
	"investool/pkg/server"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:  "investool",
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
	if err := root.Execute(); err != nil {
		log.ServiceLog().Fatal(fmt.Sprintln(os.Stderr, err))
		os.Exit(1)
	}
}

func run() {

	log.ServiceLog().Info("Start investool")
	err := config.Init("config")
	if err != nil {
		println(err)
	}
	server.Start()
}
