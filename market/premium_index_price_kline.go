package market

import "github.com/MegaBytee/bybit-go"

//docs: https://bybit-exchange.github.io/docs/v5/market/preimum-index-kline
// Limit for data size per page. [1, 1000]. Default: 200

func GetPremiumIndexPriceKline(p KlineParams) (bybit.Response, Kline) {
	bybit.Client.SetEndPoint("/v5/market/premium-index-price-kline")
	params := p.getParams()
	body, _ := bybit.Client.GetRequest(params)
	y := Kline{}
	return bybit.GetResponse(body, &y), y
}
