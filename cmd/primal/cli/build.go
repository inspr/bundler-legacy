package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/platform"
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
	platformType := args[0]
	if inputPath == "" {
		inputPath = defaultPath
	}
	if isYaml(inputPath) && validFile(inputPath) {
		var primal api.Primal
		fs := filesystem.NewMemoryFs()

		opts, err := getConfigs(inputPath)
		primal.Options = opts
		primal.Options.Root = getDirPath(inputPath)

		if !hasTemplate(primal.Options.Root) {
			fmt.Println("template file does not exists")
			return
		}

		if err != nil {
			fmt.Printf("failed to get configs from file: %v\n", err)
			return
		}

		primal.Options.Watch = false

		switch platformType {
		case api.PlatformWeb:
			primal.Options.Platform = api.PlatformWeb
		case api.PlatformElectron:
			primal.Options.Platform = api.PlatformElectron
		default:
			fmt.Printf("platform %s not supported\n", args[0])
			return
		}

		platform, err := platform.NewPlatform(primal.Options, fs)
		if err != nil {
			fmt.Printf("failed to create platform: %v\n", err)
		}

		fmt.Printf("running platform %s in build mode!\n", platformType)
		platform.Run()
		return
	}
	fmt.Print("yaml file not found for path ", inputPath, "\n")
}
