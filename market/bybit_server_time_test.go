package market

import (
	"testing"
)

func TestGetBybitServerTime(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetBybitServerTime()
			if got.TimeSecond == "" {
				t.Errorf("GetBybitServerTime() = %v, want %v", got, "timesecond")
			}
		})
	}
}
