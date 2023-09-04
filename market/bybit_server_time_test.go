package market

import (
	"reflect"
	"testing"
)

func TestGetBybitServerTime(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got, _ := GetBybitServerTime(); !reflect.DeepEqual(got.RetCode, tt.want) {
				t.Errorf("GetBybitServerTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
