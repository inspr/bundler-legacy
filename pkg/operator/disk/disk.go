package disk

import (
	"fmt"
	"os"

	"inspr.dev/primal/pkg/api"
)

type Disk struct {
	path string
	meta api.Metadata
}

func NewDisk(path string) *Disk {
	return &Disk{
		path: path,
		meta: api.NewMetadata(),
	}
}

func (h *Disk) Apply(props api.OperatorProps, opts api.OperatorOptions) {

	// loop through waiting for refresh commands
	// when op gets a signal to close then stop it
Main:
	for {
		select {
		case <-h.meta.Close:
			break Main

		case <-h.meta.Refresh:
			h.meta.Messages <- fmt.Sprintf("writing files to %s", h.path)

			if _, err := os.Stat(h.path); os.IsNotExist(err) {
				os.Mkdir(h.path, 0755)
				os.Mkdir(h.path+"/assets", 0755)
			}

			for key, file := range props.Files.Raw() {
				// TODO: catch the error and return in an "errors" chan
				f, _ := os.Create(h.path + key)
				f.Write(file)
			}

			h.meta.Done <- true
		}
	}
}

func (h *Disk) Meta() api.Metadata {
	return h.meta
}
