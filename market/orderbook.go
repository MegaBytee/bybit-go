package market

import (
	"github.com/MegaBytee/bybit-go"
)

// docs: https://bybit-exchange.github.io/docs/v5/market/orderbook
type Orderbook struct {
	S  string   `json:"s"`  //Symbol name
	B  []string `json:"b"`  //Bid, buyer. Sort by price desc
	A  []string `json:"a"`  //Ask, seller. Order by price asc
	Ts int      `json:"ts"` //The timestamp (ms) that the system generates the data
	U  int      `json:"u"`  //Update ID, is always in sequence; For future, it is corresponding to u in the wss 200-level orderbook; For spot, it is corresponding to u in the wss 50-level orderbook
}

type OrderbookParams struct {
	Category string `json:"category"` //* spot,linear,inverse,option
	Symbol   string `json:"symbol"`   //* Symbol name
	Limit    int    `json:"limit"`    //Limit size for each bid and ask ; spot: [1, 50]. Default: 1. linear&inverse: [1, 200]. Default: 25. option: [1, 25]. Default: 1.
}

func GetOrderbook(p *OrderbookParams) Orderbook {
	params := &bybit.CallParams{
		Method:   "GET",
		EndPoint: "/v5/market/orderbook",
		Params:   p,
		Fields:   []string{"Category", "Symbol"},
	}
	x := Orderbook{}
	bybit.Call(params, &x)
	return x

}
