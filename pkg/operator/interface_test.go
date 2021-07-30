package operator

import (
	"context"
	"reflect"
	"testing"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

func TestNewOperator(t *testing.T) {
	ctx := context.Background()

	options := api.PrimalOptions{}
	fs := filesystem.NewMemoryFs()

	tests := []struct {
		name string
		want *Operator
	}{
		{
			name: "new_operator",
			want: &Operator{
				Options: options,
				Ctx:     ctx,
				Fs:      fs,
			},
		},
	}

	for _, tt := range tests {
		got := NewOperator(options, fs)

		if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
			t.Errorf(
				"NewOperator() = %v, want %v",
				reflect.TypeOf(got),
				reflect.TypeOf(tt.want),
			)
		}
	}
}

func TestInitMainOperators(t *testing.T) {
	options := api.PrimalOptions{}
	fs := filesystem.NewMemoryFs()

	mainOperator := NewOperator(options, fs)

	tests := []struct {
		name       string
		init       func()
		wantLength int
	}{
		{
			name:       "main_operators_before_init",
			wantLength: 0,
		},
		{
			name: "main_operators_after_init",
			init: func() {
				mainOperator.InitMainOperators()
			},
			wantLength: 3,
		},
	}

	for _, tt := range tests {
		if tt.init != nil {
			tt.init()
		}

		gotLength := len(MainOps)

		if reflect.TypeOf(gotLength) != reflect.TypeOf(tt.wantLength) {
			t.Errorf(
				"InitMainOperators() = %v, want %v",
				reflect.TypeOf(gotLength),
				reflect.TypeOf(tt.wantLength),
			)
		}
	}
}
