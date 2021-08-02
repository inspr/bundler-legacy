package operator

import (
	"reflect"
	"testing"

	"inspr.dev/primal/pkg/api"
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

func (disk *Disk) TestTask(t *testing.T) {
	
}
