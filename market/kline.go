package market

import (
	"github.com/MegaBytee/bybit-go"
)

//docs: https://bybit-exchange.github.io/docs/v5/market/kline

type Kline struct {
	Category string   `json:"category"`
	Symbol   string   `json:"symbol"`
	List     []string `json:"list"`
}

// *: required
type KlineParams struct {
	Category string `json:"category"` //* spot,linear,inverse
	Symbol   string `json:"symbol"`   //* Symbol name
	Interval string `json:"interval"` //* Kline interval. 1,3,5,15,30,60,120,240,360,720,D,M,W
	Start    int    `json:"start"`    //The start timestamp (ms)
	End      int    `json:"end"`      //The end timestamp (ms)
	Limit    int    `json:"limit"`    //Limit for data size per page. [1, 1000]. Default: 200
}

func GetKline(p *KlineParams) (bybit.Response, Kline) {
	params := &bybit.CallParams{
		Method:   "GET",
		EndPoint: "/v5/market/kline",
		Params:   p,
		Fields:   []string{"Category", "Symbol", "Interval"},
	}
	x := Kline{}
	return bybit.Call(params, &x), x

}
