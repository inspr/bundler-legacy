package platform

import (
	"reflect"
	"testing"

	"inspr.dev/primal/pkg/api"
)

func TestPlatform_Electron(t *testing.T) {
	tests := []struct {
		name string
		p    *Platform
		want api.PlatformInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Electron(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Platform.Electron() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElectron_Run(t *testing.T) {
	tests := []struct {
		name string
		e    *Electron
	}{
		// TODO: Add test cases.
		{
			name: "",
			e:    &Electron{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Run()
		})
	}
}

func TestElectron_Watch(t *testing.T) {
	tests := []struct {
		name string
		e    *Electron
	}{
		// TODO: Add test cases.
		{
			name: "",
			e:    &Electron{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Watch()
		})
	}
}
