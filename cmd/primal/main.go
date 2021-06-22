package main

import (
	"os"

	"inspr.dev/primal/pkg/disk"
	"inspr.dev/primal/pkg/filesystem"
	"inspr.dev/primal/pkg/logger"
	"inspr.dev/primal/pkg/platform/web"
	"inspr.dev/primal/pkg/server"
)

func main() {
	root, _ := os.Getwd()
	fs := filesystem.NewMemoryFs()
	primal := NewCompiler(root, fs)

	BundlerClient := web.NewBundler().WithMinification().Target("client")
	BundlerServer := web.NewBundler().WithMinification().Target("server")

	HtmlGen := web.NewHtml()
	Disk := disk.NewDisk("./__build__")
	Server := server.NewServer(3049)
	Logger := logger.NewLogger()

	primal.
		Add(BundlerClient, BundlerServer).
		Add(HtmlGen).
		Add(Logger, Disk, Server).
		Apply()
}
