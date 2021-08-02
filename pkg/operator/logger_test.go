package operator

import (
	"reflect"
	"testing"

	"inspr.dev/primal/pkg/api"
)

func TestNewLogger(t *testing.T) {
	options := api.PrimalOptions{}
	mainOperator := NewMockOperator(options)

	tests := []struct {
		name string
		want *Logger
	}{
		{
			name: "new_html_operator",
			want: &Logger{
				Operator: &Operator{},
			},
		},
	}

	for _, tt := range tests {
		got := mainOperator.NewLogger()

		if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
			t.Errorf(
				"NewLogger() = %v, want %v",
				reflect.TypeOf(got),
				reflect.TypeOf(tt.want),
			)
		}
	}
}
