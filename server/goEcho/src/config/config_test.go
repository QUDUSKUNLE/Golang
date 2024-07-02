package config

import "testing"

func TestLoadEnvironmentVariable(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadEnvironmentVariable(); (err != nil) != tt.wantErr {
				t.Errorf("LoadEnvironmentVariable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
