package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the bulb to the desired color",
	Long:  "Sets the bulb to the desired color",
	Run: setRun,

}

func setRun (cmd *cobra.Command, args []string) {
	for _, v := range args {
		fmt.Println("setting to ", v)
	}
}

