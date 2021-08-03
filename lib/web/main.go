package main

import (
	"context"

	"inspr.dev/primal/lib/web/server"
)

func main() {

	ctx := context.Background()
	server.Start(ctx, "3058")
}
