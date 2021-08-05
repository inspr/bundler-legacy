package platform

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"inspr.dev/primal/pkg/filesystem"
)

// FileServerHandler provides the function signature for passing to the FileServerWith404
type FileServerHandler = func(w http.ResponseWriter, r *http.Request) bool

/*
FileServer wraps the http.FileServer checking to see if the url path exists in it FileSystems(filesystem.FileSystem, http.FileSystem) first.
The implementation of Handlers can choose to either modify the request, e.g. change the URL path and return true to have the
default FileServer handling to still take place, or return false to stop further processing, for example if you wanted
to write a custom response
e.g. redirects to root and continues the file serving handler chain
	func fileServerHandler404(w http.ResponseWriter, r *http.Request, ... ) (goNext bool) {
		//if not found redirect to /
		r.URL.Path = "/"
		return true
	}
Use the same as you would with a http.FileServer e.g.
	r.Handle("/", http.StripPrefix("/", FileServer(http.Dir("./staticDir"))))
*/
func FileServer(root http.FileSystem, fs filesystem.FileSystem) http.Handler {
	fileServer := http.FileServer(root)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//make sure the url path starts with /
		upath := r.URL.Path
		if !strings.HasPrefix(upath, "/") {
			upath = "/" + upath
			r.URL.Path = upath
		}
		upath = path.Clean(upath)

		// attempt to open the file via the filesystem.FileSystem
		if goNext := inMemoryHandler(w, r, fs, upath); !goNext {
			return
		}

		// attempt to open the file via the http.FileSystem
		if goNext := onDiskHandler(w, r, root, upath); !goNext {
			return
		}

		// default serve
		fileServer.ServeHTTP(w, r)
	})
}

func onDiskHandler(w http.ResponseWriter, r *http.Request, fs http.FileSystem, path string) bool {
	f, err := fs.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// go next
			w.WriteHeader(404)
			fmt.Fprintf(w, "file doesn't exist on disk. %s", err.Error())

			return false
		}
	}

	// close if successfully opened
	if err == nil {
		f.Close()
	}

	return true
}

func inMemoryHandler(w http.ResponseWriter, r *http.Request, fs filesystem.FileSystem, path string) bool {
	file, err := fs.Get(path)
	if err != nil {
		fmt.Println("err fs.get, should call next :", err)
		return true
	}

	SetContentType(w, path)
	w.Write(file)

	return false
}
