package server

import (
	"context"

	"inspr.dev/primal/lib/web/server/controller"
)

func Start(ctx context.Context, port string) {

	server := controller.NewServer(ctx, port)

	server.Run()
}
