package filesystem

// Move to interface
type FileSystem interface {
	// Get return the file in a file system, given a path (ext included)
	Get(path string) ([]byte, error)

	Write(path string, data []byte) error

	List() []string

	Raw() map[string][]byte
}
