package vm

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

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
			"test1", // Name of the TEST
			args{
				ctx:    context.Background(),
				script: "run('chico')",
			}, // Arguments for the Test
			Response{
				HTML: []byte("chico"),
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
				t.Errorf("Run result was = %v, want %v", got, tt.want)
			}

			virtual.Close()
		})
	}
}
