package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"inspr.dev/primal/lib/web/server/vm"
)

var logger *log.Logger

func init() {
	logger = log.Default()
}

var contentTypes = map[string]string{
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

// Handler is a structure which cointains methods to handle
// requests received by the UID Provider API
type Handler struct {
	ctx      context.Context
	dataPath string
}

// Handler is an alias of the api router function
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// NewHandler instantiates a new Handler structure
func NewHandler(ctx context.Context, path string) *Handler {
	return &Handler{
		ctx:      ctx,
		dataPath: path,
	}
}

// ServeFiles serves the files on the directory specified by 'InitBuildDir' handler
func (h *Handler) ServeFiles(machine vm.Interface) HandlerFunc {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: if url doesn't contain any extensions, send it to VM

		path := r.URL.Path[0:]
		if path != "/" {
			setContentType(w, path)
			path = h.dataPath + path
		} else {
			path = h.dataPath + "/index.html"
		}

		vmResponse := <-machine.Run(vm.Request{
			UUID: uuid.New(),
			Path: path,
		})

		// Return the HTML for the user, sort of
		w.Write(vmResponse.HTML)

		http.ServeFile(w, r, path)
	})
}

// HealthCheck is a handler for health checking the server
func (h *Handler) HealthCheck() HandlerFunc {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCors(w)
		writeResponse(w, http.StatusOK, "server is running!")
	})
}

// writeResponse writes a request response in its response writer
func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	setCors(w)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func setCors(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	// (*w).Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	// (*w).Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers,"+
	// 	"Authorization, X-Requested-With, Content-Length, Server, Date,"+
	// 	"Access-Control-Allow-Methods, Access-Control-Allow-Origin")
}

// setContentType adds the given file's content type to the header
func setContentType(w http.ResponseWriter, file string) {
	ext := filepath.Ext(file)
	w.Header().Add("Content-Type", contentTypes[ext])
}

// setCacheDuration adds Cache-Control header with the given amount of seconds
func setCacheDuration(w http.ResponseWriter, seconds int64) {
	w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d", seconds))
}

// validFile checks if given filePath exists
func validFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
