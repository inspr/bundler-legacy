package operator

import (
	"fmt"
	"io/ioutil"
	"strings"

	"inspr.dev/primal/pkg/filesystem"
)

// Static defines the static file generator structure
type Static struct {
	files []string
	root  string
}

// NewStatic returns a reference to a new static file generator structure
// with the given path and files
func NewStatic(root string, files []string) *Static {
	return &Static{
		files,
		root,
	}
}

// ! to be implemented
// Handler
func (s *Static) Handler(fs filesystem.FileSystem) {
	for _, path := range fs.List() {
		newPath := strings.Replace(path, ".", "", 1)
		// TODO: root
		data, err := ioutil.ReadFile(s.root + newPath)

		if err != nil {
			fmt.Println(err)
		}

		fs.Write(newPath, data)
	}
}
