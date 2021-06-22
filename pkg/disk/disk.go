package disk

import (
	"context"
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

func (h *Disk) Apply(ctx context.Context, opts api.OperatorOptions) error {
	h.meta.State <- api.WORKING
	h.meta.Progress <- 0.0

	// fmt.Println(opts.Files)

	h.meta.Messages <- fmt.Sprintf("writing files to %s", h.path)

	if _, err := os.Stat(h.path); os.IsNotExist(err) {
		os.Mkdir(h.path, 0755)
		os.Mkdir(h.path+"/assets", 0755)
	}

	for key, file := range opts.Files.Raw() {
		f, err := os.Create(h.path + key)

		if err != nil {
			return err
		}

		f.Write(file)
	}
	// opts.Files.Flush(h.path)

	h.meta.State <- api.DONE
	h.meta.Progress <- 1.0

	return nil
}

func (h *Disk) Meta() api.Metadata {
	return h.meta
}
