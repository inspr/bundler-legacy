package operator

import (
	"fmt"
	"io/ioutil"
	"strings"

	"inspr.dev/primal/pkg/api"
)

type Static struct {
	files []string
	meta  api.Metadata
}

func NewStatic(files []string) *Static {
	return &Static{
		files: files,
		meta:  api.NewMetadata(),
	}
}

func (h *Static) Apply(props api.OperatorProps, opts api.OperatorOptions) {

	var loadFiles = func() {
		for _, path := range h.files {
			newPath := strings.Replace(path, ".", "", 1)
			data, err := ioutil.ReadFile(opts.Root + newPath)

			if err != nil {
				fmt.Println(err)
			}

			props.Files.Write(newPath, data)
		}
		h.meta.Done <- true
	}

	loadFiles()
Main:
	for {
		select {
		case <-h.meta.Close:
			break Main

		case <-h.meta.Refresh:
			loadFiles()
		}
	}

}

func (h *Static) Meta() api.Metadata {
	return h.meta
}
