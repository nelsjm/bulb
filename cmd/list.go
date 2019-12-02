package cmd

import (
	"fmt"
	"os"

	"github.com/nelsjm/bulb/internal/settings"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of the available commands",
	Long:  "Lists all of the available commads from the settings",
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	cmds, err := settings.GetSettings()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	count := 0
	for k, v := range cmds {
		count++
		fmt.Printf("%v) %s - %s\n", count, k, v)
	}
}
