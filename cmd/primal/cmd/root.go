package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

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

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".primal")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
