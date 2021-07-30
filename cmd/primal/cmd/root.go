package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "primal [mode] [platform]",
	Short: "Primal CLI",
	Long:  `A framework for developing web applications`,
	Example: "primal develop web -f dir/config.yaml\n" +
		"primal build electron -f dir/config.yaml\n",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("main command of the primal cli, to see the full list of " +
			"existent subcommands please use 'primal help'")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd.PersistentFlags().StringP("file", "f", inputPath, "-f path/to/file.yaml")

	rootCmd.AddCommand(developCmd, buildCmd)

	rootCmd.Execute()
}
