package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/platform"
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

		if !hasTemplateFolder(primal.Options.Root) {
			fmt.Println("template folder does not exist")
			return
		}

		if err != nil {
			fmt.Printf("failed to get configs from file: %v\n", err)
			return
		}

		primal.Options.Watch = true

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

		fmt.Printf("running platform %s in development mode!\n", platformType)

		ctx, cancel := context.WithCancel(context.Background())
		platform.Watch(ctx, cancel)
		return
	}
	fmt.Print("yaml file not found for path ", inputPath, "\n")
}
