package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		config      *Conf
		wantConsole bool
		wantRemote  bool
	}{
		{
			name: "console only",
			config: &Conf{
				Console: true,
				Remote:  false,
			},
			wantConsole: true,
			wantRemote:  false,
		},
		{
			name: "remote only",
			config: &Conf{
				Console: false,
				Remote:  true,
			},
			wantConsole: false,
			wantRemote:  true,
		},
		{
			name: "both console and remote",
			config: &Conf{
				Console: true,
				Remote:  true,
			},
			wantConsole: true,
			wantRemote:  true,
		},
		{
			name: "neither console nor remote",
			config: &Conf{
				Console: false,
				Remote:  false,
			},
			wantConsole: false,
			wantRemote:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleCalled := false
			handle := func(b []byte) {
				handleCalled = true
			}

			logger := New(tt.config, handle)

			logger.Info("测试", zap.Bool("cancel", handleCalled))
		})
	}
}
