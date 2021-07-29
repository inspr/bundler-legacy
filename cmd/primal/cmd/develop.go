package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// developCmd represents the develop command
var developCmd = &cobra.Command{
	Use:   "develop",
	Short: "Starts the compiler in watch mode",
	Long: `Starts the compiler in watch mode. It doesn’t minify the code.
	
	The platform will be one of “electron”, “native”, “web”`,

	Run: func(cmd *cobra.Command, args []string) {
		fpath, _ = cmd.Flags().GetString("fpath")
		develop(args)
	},
}

func init() {
	rootCmd.AddCommand(developCmd)
	// developCmd.PersistentFlags().String("foo", "", "A help for foo")

	currPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	Path = currPath + "/config.yaml" // ou yml

	developCmd.Flags().StringP("fpath", "f", Path, "add path")
}

func develop(args []string) {
	if fpath == "" {
		fpath = Path
	}
	if isYaml() {

		if args[0] == "electron" {
			fmt.Print("develop in watch mode on ", args[0], " in ", fpath, "\n")
			readFile()
			return
		}
		if args[0] == "web" {
			fmt.Print("develop in watch mode on ", args[0], " in ", fpath, "\n")
			readFile()
			return
		}
		if args[0] == "native" {
			fmt.Print("develop in watch mode on ", args[0], " in ", fpath, "\n")
			readFile()
			return
		}
		fmt.Print("unexisting platform\n")
	}
	fmt.Print("yaml file not found\n")
}
