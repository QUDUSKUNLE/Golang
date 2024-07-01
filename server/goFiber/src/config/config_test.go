package config

import "testing"

func TestConfig(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Config(tt.args.key); got != tt.want {
				t.Errorf("Config() = %v, want %v", got, tt.want)
			}
		})
	}
}
