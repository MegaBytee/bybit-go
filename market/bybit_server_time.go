package market

import "github.com/MegaBytee/bybit-go"

//docs: https://bybit-exchange.github.io/docs/v5/market/time

type ServerTime struct {
	TimeSecond string `json:"timeSecond"`
	TimeNano   string `json:"timeNano"`
}

func GetBybitServerTime() (bybit.Response, ServerTime) {
	bybit.Client.SetEndPoint("/v5/market/time")
	body, _ := bybit.Client.GetRequest("")
	y := ServerTime{}
	return bybit.GetResponse(body, &y), y
}
