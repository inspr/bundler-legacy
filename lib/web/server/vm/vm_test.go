package vm

import (
	"context"
	"reflect"
	"testing"

	_ "embed"

	"github.com/google/uuid"
)

// go:embed mock_entry
var embbededScript string

func TestNewVM(t *testing.T) {
	type args struct {
		ctx    context.Context
		script string
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "run VM with valid script content",
			args: args{
				ctx:    context.Background(),
				script: embbededScript,
			}, // Arguments for the Test
			want: Response{
				HTML: []byte("oi"),
			},
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			virtual := New(tt.args.ctx)
			virtual.WithScript(tt.args.script)

			result := virtual.Run(Request{
				UUID: uuid.New(),
			})

			if got := <-result; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run result was = %v, want %v", string(got.HTML), tt.want)
			}

			virtual.Close()
		})
	}
}
