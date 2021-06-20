package disk

import (
	"context"
	"fmt"

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
	opts.Files.Flush(h.path)

	h.meta.State <- api.DONE
	h.meta.Progress <- 1.0

	return nil
}

func (h *Disk) Meta() api.Metadata {
	return h.meta
}
