package operator

import (
	"os"
	"path"
	"reflect"
	"testing"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/workflow"
)

func TestNewDisk(t *testing.T) {
	options := api.PrimalOptions{}
	mainOperator := NewMockOperator(options)

	tests := []struct {
		name string
		want *Disk
	}{
		{
			name: "new_disk_operator",
			want: &Disk{
				Operator: &Operator{},
			},
		},
	}

	for _, tt := range tests {
		got := mainOperator.NewDisk()

		if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
			t.Errorf(
				"NewDisk() = %v, want %v",
				reflect.TypeOf(got),
				reflect.TypeOf(tt.want),
			)
		}
	}
}

func TestDisk_Task(t *testing.T) {
	currDir, _ := os.Getwd()
	root := path.Join(currDir, "../../")
	buildFolder := path.Join(root, "__build__")

	diskOperator := NewMockOperator(api.PrimalOptions{
		Root: root,
	}).NewDisk()

	testWorkflow := workflow.NewWorkflow()
	testWorkflow.Add(diskOperator.Task())

	tests := []struct {
		name        string
		fileName    string
		fileContent []byte
		wantErr     bool
	}{
		{
			name:        "write_file_to_disk",
			fileName:    "/test.txt",
			fileContent: []byte("test text"),
			wantErr:     false,
		},
		{
			name:        "failes_to_write_file_to_disk",
			fileName:    "/no-folder/test.txt",
			fileContent: []byte("test text"),
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		diskOperator.Fs.Write(tt.fileName, tt.fileContent)

		testWorkflow.Run()

		if _, err := os.Stat(buildFolder + tt.fileName); os.IsNotExist(err) {
			if !tt.wantErr {
				t.Errorf("Disk.TestTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		}
	}

	os.RemoveAll(buildFolder)
}
