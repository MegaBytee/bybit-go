package market

import "github.com/MegaBytee/bybit-go"

//docs: https://bybit-exchange.github.io/docs/v5/market/index-kline
// Limit for data size per page. [1, 1000]. Default: 200

func GetIndexPriceKline(p *KlineParams) (bybit.Response, Kline) {
	params := &bybit.CallParams{
		Method:   "GET",
		EndPoint: "/v5/market/index-price-kline",
		Params:   p,
		Fields:   []string{"Category", "Symbol", "Interval"},
	}
	x := Kline{}
	return bybit.Call(params, &x), x

}
