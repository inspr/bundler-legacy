package operator

import (
	"fmt"

	"inspr.dev/primal/pkg/filesystem"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (*Logger) Handler(fs filesystem.FileSystem) {
	fmt.Println(fs)
}
