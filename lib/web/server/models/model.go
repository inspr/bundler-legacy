package models

// ServerDI structure contains a path to the directory
// which contains the files to be served
type ServerDI struct {
	Path string `yaml:"path"`
}
