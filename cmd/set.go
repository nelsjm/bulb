package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nelsjm/bulb/internal/settings"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the bulb to the desired color",
	Long:  "Sets the bulb to the desired color",
	Run:   setRun,
}

func setRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("The set command expects exactly 1 argument")
		os.Exit(1)
	}

	clr := args[0]

	url, err := settings.GetURL(clr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if url == "" {
		fmt.Println("No setting found for", clr)
		os.Exit(1)
	}

	cli := http.Client{Timeout: 30 * time.Second}
	resp, err := cli.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Got a status code of %v instead of a 200", resp.StatusCode)
		os.Exit(1)
	}

	fmt.Println("Status set to", clr)
}
