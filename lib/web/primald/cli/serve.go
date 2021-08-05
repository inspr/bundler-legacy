package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"inspr.dev/primal/lib/web/server/controller"
	"inspr.dev/primal/lib/web/server/vm"
)

// serveCmd implements primald serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs Primald to serve compiled files",
	Long: "Runs Primald to serve compiled files from the specified " +
		"directory in the specified port. If not specified, the default dir " +
		"is the current directory, and the default port is 3058",
	Example: "primal serve -f dir/__build__ -p 8080",

	Run: func(cmd *cobra.Command, args []string) {
		inputPath, _ = cmd.Flags().GetString("file")
		inputPort, _ = cmd.Flags().GetString("port")
		serve(args)
	},
}

func init() {
	serveCmd.Flags().StringP("file", "f", inputPath, "-f path/to/dir/__build__")
	serveCmd.Flags().StringP("port", "p", inputPort, "-p 3058")
}

func serve(args []string) {
	var err error

	if inputPort == "" {
		inputPort = "3058"
	}

	if inputPath == "" {
		inputPath, err = os.Getwd()
		if err != nil {
			panic(err)
		}
		inputPath += "/__build__"
	}

	if validFile(inputPath) {
		ctx := context.Background()

		fmt.Printf("serving files in %s on port %s\n", inputPath, inputPort)

		// TODO: Instantiate VM and send it's reference to Server

		machine := vm.New(ctx)
		machine.WithScript("run('boy')")

		server := controller.NewServer(ctx, inputPort, inputPath, machine)
		server.Run()
		return
	}

	fmt.Println("path '", inputPath, "' not found")
}
