package platform

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"golang.org/x/net/websocket"
	"inspr.dev/primal/pkg/filesystem"
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

// TODO:
// T is a test struct
type T struct {
	Msg  string
	Path string
}

// jsonRecieveServer prints message sent from client using websocket.JSON
func jsonRecieveServer(ws *websocket.Conn) {
	for {
		var msg T
		// Receive receives a text message serialized T as JSON.
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recived:%#v\n", msg)

		defer ws.Close()
	}
}

// updateServer listens for server reload chanel and sends to client message about update using websocket.JSON
func (s *Server) updateServer(ws *websocket.Conn) {
	for {
		<-(s.reload)
		msg := "bundle updated"
		err := websocket.JSON.Send(ws, msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%#v\n", msg)

		defer ws.Close()
	}
}

// Start runs the dev server with the given filesystem in localhost:3049
func (s *Server) Start(files filesystem.FileSystem) {
	// WS handlers
	http.Handle("/json", websocket.Handler(jsonRecieveServer))
	http.Handle("/got-update", websocket.Handler(s.updateServer))

	go http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var file []byte
		var err error

		path := r.URL.Path[0:]

		switch path {
		case "/":
			file, err = files.Get("/index.html")
		default:
			file, err = files.Get(path)
		}

		if err == nil {
			SetContentType(w, path)
			// SetCacheDuration(w, 31536000)
			w.Write(file)
		} else {
			w.WriteHeader(404)
		}
	})

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
