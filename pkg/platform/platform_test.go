package platform

import (
	"testing"

	"inspr.dev/primal/pkg/api"
	"inspr.dev/primal/pkg/filesystem"
)

func TestNewPlatform(t *testing.T) {
	type args struct {
		options api.PrimalOptions
		fs      filesystem.FileSystem
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "generates web platform",
			args: args{
				fs: filesystem.NewMemoryFs(),
				options: api.PrimalOptions{
					Root:     "/random/path",
					Platform: api.PlatformWeb,
				},
			},
		},
		{
			name: "generates electron platform",
			args: args{
				fs: filesystem.NewMemoryFs(),
				options: api.PrimalOptions{
					Root:     "/random/path",
					Platform: api.PlatformElectron,
				},
			},
		},
		{
			name: "tries to create invalid platform",
			args: args{
				fs: filesystem.NewMemoryFs(),
				options: api.PrimalOptions{
					Root:     "/random/path",
					Platform: "invalid",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlatform(tt.args.options, tt.args.fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlatform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && !tt.wantErr {
				t.Errorf("NewPlatform() didn't return a valid platform")
			}
		})
	}
}
