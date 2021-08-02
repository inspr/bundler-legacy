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
	ctx context.Context
}

// Handler is an alias of the api router function
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// NewHandler instantiates a new Handler structure
func NewHandler(ctx context.Context) *Handler {
	return &Handler{
		ctx: ctx,
	}
}

func (h *Handler) ServerHandler() HandlerFunc {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data models.ServerDI
		path := r.URL.Path[0:]
		setContentType(w, path)

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response := fmt.Sprintf("error while decoding request body: %v", err)
			writeResponse(w, http.StatusBadRequest, response)
			return
		}

		if validFile(data.Path) {
			// TODO: get files and use the server to serve them
		}

		logger.Printf("file %s not found", path)
		writeResponse(w, http.StatusNotFound, "file not found for given path")
		return
	})
}

func (h *Handler) HealthCheck() HandlerFunc {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCors(w)
		fmt.Println("health checking")
		writeResponse(w, http.StatusOK, "server is")
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
func validFile(filePath string) bool {
	var err error
	if _, err = os.Stat(filePath); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	panic(err)
}
