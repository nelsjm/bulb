package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:"add",
	Short: "Adds a new setting to the config",
	Long: "Adds a new setting to the config",
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	for _, v := range args {
		fmt.Println("Adding with", v)
	}
}
