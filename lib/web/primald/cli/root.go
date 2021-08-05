package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "primald serve",
	Short:   "Primal Daemon CLI",
	Long:    "Server application that gets compiled files and serves them",
	Example: "primald serve -f dir/__build__ -p 8080",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("main command of the primald cli, to see the full list of " +
			"existent subcommands please use 'primald help'")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd.AddCommand(serveCmd)

	rootCmd.Execute()
}
