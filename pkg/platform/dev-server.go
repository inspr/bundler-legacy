package platform

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"golang.org/x/net/websocket"
	"inspr.dev/primal/pkg/filesystem"
)

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

func SetContentType(w http.ResponseWriter, file string) {
	ext := filepath.Ext(file)
	w.Header().Add("Content-Type", ContentTypes[ext])
}

func SetCacheDuration(w http.ResponseWriter, seconds int64) {
	w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d", seconds))
}

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	for {
		<-time.After(1000 * time.Millisecond)
		ws.Write([]byte("sdss"))
	}
}

func Start(files filesystem.FileSystem) {
	fmt.Printf("Available on http://127.0.0.1:%d\n", 3049)

	go http.Handle("/ws", websocket.Handler(EchoServer))

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
			SetCacheDuration(w, 31536000)
			w.Write(file)
		} else {
			w.WriteHeader(404)
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 3049), nil))
}
