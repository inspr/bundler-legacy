package operator

import (
	"context"
	"os"
	"reflect"
	"testing"

	"inspr.dev/primal/pkg/api"
)

func TestNewStatic(t *testing.T) {
	options := api.PrimalOptions{}
	mainOperator := NewMockOperator(options)

	testFiles := []string{"sw.js"}

	tests := []struct {
		name string
		want *Static
	}{
		{
			name: "new_html_operator",
			want: &Static{
				Operator: &Operator{},
				files:    testFiles,
			},
		},
	}

	for _, tt := range tests {
		got := mainOperator.NewStatic(testFiles)

		if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
			t.Errorf(
				"NewLogger() = %v, want %v",
				reflect.TypeOf(got),
				reflect.TypeOf(tt.want),
			)
		}
	}
}

func TestStatic_Task(t *testing.T) {
	path, _ := os.Getwd()

	options := api.PrimalOptions{
		Root: path,
	}

	tests := []struct {
		name         string
		filesToWrite []string
		wantErr      bool
	}{
		{
			name:         "static_adds_file_to_build",
			filesToWrite: []string{"static.go"},
			wantErr:      false,
		},
		// TODO: Return Error from workflow.Task.Run method to handle errors
		// {
		// 	name:         "static_fails_to_add_file_doesn't_exist",
		// 	filesToWrite: []string{"non-existent.js"},
		// 	wantErr:      true,
		// },
	}

	for _, tt := range tests {
		staticOperator := NewMockOperator(options).NewStatic(tt.filesToWrite)
		got := staticOperator.Task()

		got.Run(context.Background(), &got)

		for _, file := range tt.filesToWrite {
			_, err := staticOperator.Fs.Get("/" + file)

			if (err != nil) != tt.wantErr {
				t.Errorf("Static.TestTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		}

	}
}
