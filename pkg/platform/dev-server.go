package platform

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"inspr.dev/primal/pkg/filesystem"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Server struct {
	reload (chan bool)
}

// Start runs the dev server with the given filesystem in localhost:3049
func (s *Server) Start(files filesystem.FileSystem) {
	// HotReload
	go http.HandleFunc("/hmr", s.SendBundleUpdates)

	// FileServer
	fileServer := FileServer(http.Dir("__build__"), files)
	go http.Handle("/", http.StripPrefix("/", fileServer))

	fmt.Printf("Available on http://127.0.0.1:%d\n", 3049)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 3049), nil))
}

// GracefullShutdown gracefully shuts down the platform bein executed
func GracefullShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	os.Exit(1)
}

type UpdateMessage struct {
	Updated bool
	Errors  bool
}

// SendBundleUpdates handler is write only WebSocket connection
func (s *Server) SendBundleUpdates(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "websocket is closed")

	ctx, cancel := context.WithTimeout(r.Context(), time.Minute*10)
	defer cancel()

	ctx = c.CloseRead(ctx)

	for {
		select {
		case <-ctx.Done():
			c.Close(websocket.StatusNormalClosure, "")
			return
		case <-(s.reload):
			msg := UpdateMessage{
				Updated: true,
				Errors:  false,
			}
			err = wsjson.Write(ctx, c, msg)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
