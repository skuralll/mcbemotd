package mcbemotd

import (
	"testing"
)

func TestGetServerInfo(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{"LocalHost", args{"localhost:19132"}},
		{"BadIP", args{"10000.00.10.1:19132"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetServerInfo(tt.args.address)
			if err != nil {
				t.Log("\nError:\n%w\n", err)
			} else {
				t.Logf("\nSuccess:\n%v\n", got)
			}
		})
	}
}
