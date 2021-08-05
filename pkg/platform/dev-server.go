package platform

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"inspr.dev/primal/pkg/filesystem"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Server struct {
	reload (chan bool)
}

// ContentTypes defines the supported file content types
var ContentTypes = map[string]string{
	".css": "text/css; charset=UTF-8",

	".js":  "application/javascript; charset=UTF-8",
	".mjs": "application/javascript; charset=UTF-8",

	".json":   "application/json; charset=UTF-8",
	".jsonld": "application/ld+json; charset=UTF-8",

	".png":  "image/png",
	".webp": "image/webp",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".svg":  "image/svg+xml; charset=utf-8",

	".woff":  "font/woff",
	".woff2": "font/woff2",
}

// SetContentType adds the given file's content type to the header
func SetContentType(w http.ResponseWriter, file string) {
	ext := filepath.Ext(file)
	w.Header().Add("Content-Type", ContentTypes[ext])
}

// SetCacheDuration adds Cache-Control header with the given amount of seconds
func SetCacheDuration(w http.ResponseWriter, seconds int64) {
	w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d", seconds))
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

// Start runs the dev server with the given filesystem in localhost:3049
func (s *Server) Start(ctx context.Context, files filesystem.FileSystem) {
	// HotReload
	go http.HandleFunc("/hmr", s.SendBundleUpdates)

	server := &http.Server{Addr: ":8080", Handler: nil}

	fmt.Printf("Available on http://127.0.0.1:%d\n", 3049)

	go server.ListenAndServe()

	<-ctx.Done()

	GracefullShutdown(server)
}

// GracefullShutdown gracefully shuts down the platform bein executed
func GracefullShutdown(server *http.Server) {
	fmt.Println("gracefully shutting down dev server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic("")
	}
	os.Exit(1)
}
