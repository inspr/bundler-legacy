package operator

import (
	"reflect"
	"testing"

	"inspr.dev/primal/pkg/api"
)

func TestNewHtml(t *testing.T) {
	options := api.PrimalOptions{}
	mainOperator := NewMockOperator(options)

	tests := []struct {
		name string
		want *Html
	}{
		{
			name: "new_html_operator",
			want: &Html{
				Operator: &Operator{},
			},
		},
	}

	for _, tt := range tests {
		got := mainOperator.NewHtml()

		if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
			t.Errorf(
				"NewHtml() = %v, want %v",
				reflect.TypeOf(got),
				reflect.TypeOf(tt.want),
			)
		}
	}
}

func TestHtml_Task(t *testing.T) {
	options := api.PrimalOptions{}
	htmlOperator := NewMockOperator(options).NewHtml()

	tests := []struct {
		name         string
		wantFileName string
		wantErr      bool
	}{
		{
			name:         "write_html_file_to_fs",
			wantFileName: "/index.html",
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		got := htmlOperator.Task()

		got.Run(&got)
		_, err := htmlOperator.Fs.Get(tt.wantFileName)

		if (err != nil) != tt.wantErr {
			t.Errorf("Html.TestTask() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}
