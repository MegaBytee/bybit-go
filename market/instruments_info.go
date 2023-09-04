package market

import (
	"github.com/MegaBytee/bybit-go"
)

// docs: https://bybit-exchange.github.io/docs/v5/market/instrument
type InstrumentParams struct {
	Category string `json:"category"` //* Product type. spot,linear,inverse,option
	Symbol   string `json:"symbol"`   //Symbol name
	Status   string `json:"status "`  //Symbol status filter spot/linear/inverse has Trading only
	BaseCoin string `json:"baseCoin"` //Base coin. linear,inverse,option only For option, it returns BTC by default
	Limit    int    `json:"limit"`    //Limit for data size per page. [1, 1000]. Default: 500
	Cursor   string `json:"cursor"`   // Use the nextPageCursor token from the response to retrieve the next page of the result set
}

type InstrumentsInfo struct {
	Category       string   `json:"category"`
	List           []string `json:"list"`
	NextPageCursor string   `json:"nextPageCursor"`
}

func GetInstrumentsInfo(p *InstrumentParams) (bybit.Response, InstrumentsInfo) {
	bybit.Client.SetEndPoint("/v5/market/instruments-info")
	y := InstrumentsInfo{}
	required_fields := []string{"Category"}
	return bybit.Execute("GET", p, required_fields, &y), y
}
