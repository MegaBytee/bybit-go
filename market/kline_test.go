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
	type args struct {
		p *KlineParams
	}
	tests := []struct {
		name       string
		args       args
		wantCode   int
		wantSymbol string
	}{
		{
			name: "testing_get_kline",
			args: args{
				p: &correct_params,
			},
			wantCode:   0,
			wantSymbol: correct_params.Symbol,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetKline(tt.args.p)
			if !reflect.DeepEqual(got.RetCode, tt.wantCode) {
				t.Errorf("GetKline() got = %v, want %v", got, tt.wantCode)
			}
			if !reflect.DeepEqual(got1.Symbol, tt.wantSymbol) {
				t.Errorf("GetKline() got1 = %v, want %v", got1, tt.wantSymbol)
			}
		})
	}
}
