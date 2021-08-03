package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"inspr.dev/primal/lib/web/server/models"
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
	DataPath string
}

// Handler is an alias of the api router function
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// NewHandler instantiates a new Handler structure
func NewHandler(ctx context.Context) *Handler {
	return &Handler{
		ctx: ctx,
	}
}

// InitBuildDir handles request that defines the directory which contains the files to be served
func (h *Handler) InitBuildDir() HandlerFunc {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data models.ServerDI

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response := fmt.Sprintf("error while decoding request body: %v", err)
			writeResponse(w, http.StatusBadRequest, response)
			return
		}

		if validFile(data.Path) {
			h.DataPath = data.Path
			return
		}

		logger.Printf("file %s not found", data.Path)
		writeResponse(w, http.StatusNotFound, "file not found for given path")
	})
}

// ServeFiles serves the files on the directory specified by 'InitBuildDir' handler
func (h *Handler) ServeFiles() HandlerFunc {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path[0:]
		if path != "/" {
			setContentType(w, path)
			path = h.DataPath + path
		} else {
			path = h.DataPath + "/index.html"
		}

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
