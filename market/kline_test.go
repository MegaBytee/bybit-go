package market

import (
	"reflect"
	"testing"
)

// get params

func TestGetKline(t *testing.T) {
	correct_params := KlineParams{
		Category: "spot",
		Symbol:   "ETHUSDT",
		Interval: "1",
		Limit:    10,
	}

	empty_params := KlineParams{}
	type args struct {
		p *KlineParams
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "correct_params_get_kline",
			args: args{
				p: &correct_params,
			},
			want: correct_params.Category,
		},
		{
			name: "empty_params_get_kline",
			args: args{
				p: &empty_params,
			},
			want: empty_params.Category,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetKline(tt.args.p)
			if !reflect.DeepEqual(got.Category, tt.want) {
				t.Errorf("GetKline() got = %v, want %v", got, tt.want)
			}

		})
	}
}
