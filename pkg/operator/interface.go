package operator

import "inspr.dev/primal/pkg/filesystem"

type Operator interface {
	Handler(fs filesystem.FileSystem)
}
