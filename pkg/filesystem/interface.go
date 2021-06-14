package filesystem

// Move to interface
type FileSystem interface {

	// Get return the file in a file system, given a path (ext included)
	Get(path string) ([]byte, error)

	// Match return a list of file paths that match the requested regex
	// this function is useful to select files to apply operations
	Match(regex string) ([]string, error)

	// GetAll return the files that match a given regex
	// think of it as a combination of Get and Match
	GetAll(regex string) ([]byte, error)

	Write(path string, data []byte) error

	List() []string

	Flush(path string) error
}
