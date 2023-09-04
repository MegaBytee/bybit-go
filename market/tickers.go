package market

import "github.com/MegaBytee/bybit-go"

// docs: https://bybit-exchange.github.io/docs/v5/market/tickers

type Tickers struct {
	Category string   `json:"category"`
	List     []string `json:"list"`
}

type TickerParams struct {
	Category string `json:"category"` //* Product type. spot,linear,inverse,option
	Symbol   string `json:"symbol"`   // Symbol name
	BaseCoin string `json:"baseCoin"` //Base coin. For option only
	ExpDate  string `json:"expDate"`  //Expiry date. e.g., 25DEC22. For option only
}

func GetTickers(p *TickerParams) (bybit.Response, Tickers) {
	bybit.Client.SetEndPoint("/v5/market/tickers")
	y := Tickers{}
	required_fields := []string{"Category"}
	return bybit.Execute("GET", p, required_fields, &y), y
}
