package market

import "github.com/MegaBytee/bybit-go"

//docs: https://bybit-exchange.github.io/docs/v5/market/time

type ServerTime struct {
	TimeSecond string `json:"timeSecond"`
	TimeNano   string `json:"timeNano"`
}

func GetBybitServerTime() (bybit.Response, ServerTime) {
	params := &bybit.CallParams{
		Method:   "GET",
		EndPoint: "/v5/market/time",
	}
	x := ServerTime{}
	return bybit.Call(params, &x), x
}
