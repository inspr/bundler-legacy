package cli

import (
	"os"
)

var inputPort string
var inputPath string

func validFile(filePath string) bool {
	var err error
	if _, err = os.Stat(filePath); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	panic(err)
}
