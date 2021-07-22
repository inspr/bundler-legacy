package operator

import (
	"fmt"
	"io/ioutil"
	"strings"

	"inspr.dev/primal/pkg/filesystem"
)

type Static struct {
	files []string
	root  string
}

func NewStatic(root string, files []string) *Static {
	return &Static{
		files,
		root,
	}
}

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
