package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Path adn f path imported from develop.go
// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the files and stops.",
	Long: `Builds the files and stops. Doesn’t run the server. Doesn’t have splitting nor ESModules.

	Builds entry-server on top of entry-browser.`,
	Run: func(cmd *cobra.Command, args []string) {
		fpath, _ = cmd.Flags().GetString("fpath")
		build(args)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	currPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	Path = currPath + "/config.yml" // ou yml

	buildCmd.Flags().StringP("fpath", "f", Path, "add path")
}

func build(args []string) {
	if fpath == "" {
		fpath = Path
	}
	if isYaml() && exists() { //imported from develop.go

		if args[0] == "electron" {
			fmt.Print("build file ", fpath, " on ", args[0], "\n")
			readFile()
			return
		}
		if args[0] == "web" {
			fmt.Print("build file ", fpath, " on ", args[0], "\n")
			readFile()
			return
		}
		if args[0] == "native" {
			fmt.Print("build file ", fpath, " on ", args[0], "\n")
			readFile()
			return
		}
		fmt.Print("unexisting platform\n")
	}
	fmt.Print("yaml file not found\n")
}
