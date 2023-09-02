package market

import (
	"log"
	"reflect"
	"testing"
)

// get params
func TestKlineParams(t *testing.T) {
	empty_params := KlineParams{}
	correct_params := KlineParams{
		Category: "spot",
		Symbol:   "ETHUSD",
		Interval: "1",
		Limit:    100,
	}
	missing_parrams := KlineParams{
		Category: "spot",
	}

	tests := []struct {
		name string
		p    *KlineParams
		want string
	}{
		{
			name: "empty_testing_kline_params",
			p:    &empty_params,
			want: empty_params.getParams(),
		},
		{
			name: "correct_testing_kline_params",
			p:    &correct_params,
			want: correct_params.getParams(),
		},
		{
			name: "missing_testing_kline_params",
			p:    &missing_parrams,
			want: missing_parrams.getParams(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.getParams(); got != tt.want {
				t.Errorf("KlineParams.getParams() = %v, want %v", got, tt.want)
			} else {
				log.Default().Println(got)
			}
		})
	}
}

func TestGetKline(t *testing.T) {
	correct_params := KlineParams{
		Category: "spot",
		Symbol:   "ETHUSDT",
		Interval: "1",
		Limit:    10,
	}
	type args struct {
		p KlineParams
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
				p: correct_params,
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
