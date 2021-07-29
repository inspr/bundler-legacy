package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// buildCmd implements primal build command
var buildCmd = &cobra.Command{
	Use:   "build [platform]",
	Short: "Runs Primal in production mode",
	Long: `Runs Primal in production mode, building and generating production-ready files.
	The platform can be one of “electron”, “web”`,
	Example: "primal build electron -f dir/config.yaml",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		inputPath, _ = cmd.Flags().GetString("file")
		runBuild(args)
	},
}

func init() {
	currPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	defaultPath = currPath + "/config.yaml"
}

func runBuild(args []string) {
	if inputPath == "" {
		inputPath = defaultPath
	}
	if isYaml(inputPath) && validFile(inputPath) {

		switch args[0] {
		case "web":
			fmt.Print("build mode on ", args[0], " in ", inputPath, "\n")
			readFile()
		case "electron":
			fmt.Print("build mode on ", args[0], " in ", inputPath, "\n")
			readFile()
		default:
			fmt.Printf("platform %s not supported\n", args[0])
		}
		return
	}
	fmt.Print("yaml file not found for path ", inputPath, "\n")
}
