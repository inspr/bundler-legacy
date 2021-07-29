package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// developCmd implements primal develop command
var developCmd = &cobra.Command{
	Use:   "develop [platform]",
	Short: "Runs Primal in watch mode",
	Long: `Runs Primal in watch mode, so changes made are automatically applied.
	The platform can be one of “electron”, “web”`,
	Example: "primal develop web -f dir/config.yaml",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		inputPath, _ = cmd.Flags().GetString("file")
		runDevelop(args)
	},
}

func init() {
	currPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	defaultPath = currPath + "/config.yaml"
}

func runDevelop(args []string) {
	if inputPath == "" {
		inputPath = defaultPath
	}
	if isYaml(inputPath) && validFile(inputPath) {

		switch args[0] {
		case "web":
			fmt.Print("develop mode on ", args[0], " in ", inputPath, "\n")
			readFile()
		case "electron":
			fmt.Print("develop mode on ", args[0], " in ", inputPath, "\n")
			readFile()
		default:
			fmt.Printf("platform %s not supported\n", args[0])
		}
		return
	}
	fmt.Print("yaml file not found for path ", inputPath, "\n")
}
