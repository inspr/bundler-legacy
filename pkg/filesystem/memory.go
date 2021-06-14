package filesystem

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// // Move to interface
// type FileSystem interface {

// 	// Get return the file in a file system, given a path (ext included)
// 	Get(path string) ([]byte, error)

// 	// Match return a list of file paths that match the requested regex
// 	// this function is useful to select files to apply operations
// 	Match(regex string) ([]string, error)

// 	// GetAll return the files that match a given regex
// 	// think of it as a combination of Get and Match
// 	GetAll(regex string) ([]byte, error)

// 	WriteFile(path string, data []byte)
// }

type MemoryFs struct {
	store map[string][]byte
}

func NewMemoryFs() *MemoryFs {
	return &MemoryFs{
		store: make(map[string][]byte),
	}
}

func (mfs *MemoryFs) Get(path string) ([]byte, error) {
	data, ok := mfs.store[path]

	if ok {
		return data, nil
	} else {
		return nil, fmt.Errorf("file %s doesn't exist in filesystem", path)
	}
}

func (mfs *MemoryFs) Match(regex string) ([]string, error) {
	return nil, nil
}

func (mfs *MemoryFs) GetAll(regex string) ([]byte, error) {
	return nil, nil
}

func (mfs *MemoryFs) Write(path string, data []byte) error {
	_, ok := mfs.store[path]

	if ok {
		return errors.New("file already exists in memory")
	}

	mfs.store[path] = data

	return nil
}

func (mfs *MemoryFs) List() []string {
	paths := []string{}

	for key := range mfs.store {
		paths = append(paths, key)
	}

	return paths
}

func ByteCountSI(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func (mfs *MemoryFs) String() string {
	writer := bytes.NewBufferString("")

	data := [][]string{}

	for key, value := range mfs.store {
		data = append(data, []string{key, ByteCountSI(len(value))})
	}

	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"file", "size"})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgGreenColor},
		tablewriter.Colors{},
	)

	table.AppendBulk(data)
	table.Render()

	return writer.String()
}

func (mfs *MemoryFs) Flush(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		os.Mkdir(path+"/assets", 0755)
	}

	for key, file := range mfs.store {
		f, err := os.Create(path + key)

		if err != nil {
			return err
		}

		f.Write(file)
	}

	return nil
}
