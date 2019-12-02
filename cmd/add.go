package cmd

import (
	"fmt"
	"os"

	"github.com/nelsjm/bulb/internal/settings"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new setting to the config",
	Long:  "Adds a new setting to the config",
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("The add command expects exactly 2 arguments")
		os.Exit(1)
	}

	clr := args[0]
	url := args[1]

	err := settings.AddURL(clr, url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Added %s at %s", clr, url)
}
