<p align="center">
  <img height="400" src="/gopher.png">
  <h1 align="center">
    Bybit-Go
    <br>
    <a href="https://megabytee.com/"><img alt="Static Badge" src="https://img.shields.io/badge/(c)-MegaBytee.com-blue"></a>
    <a href="https://github.com/MegaBytee/bybit-go/blob/master/LICENSE" ><img alt="Static Badge" src="https://img.shields.io/badge/license-MIT-red">
    </a>
  </h1>
</p>


# bybit-go
Go API V5 connector for Bybit's HTTP and WebSockets APIs.

<a href="https://bybit-exchange.github.io/docs/v5/intro">bybit docs</a>

//not finished yet still working on this...
//TODO description and how it work 

## Markets

| Endpoint | Description | Status |
|:-------:|:----------- |:------:|
| [Get Bybit Server Time](/market/bybit_server_time.go) | GET /v5/market/time | ✔ |
| [Get Kline](/market/kline.go) | GET /v5/market/kline | ✔ |
| [Get Mark Price Kline](/market/mark_price_kline.go) | GET /v5/market/mark-price-kline | ✔ |
| [Get Index Price Kline](/market/mark_index_kline.go) | GET /v5/market/index-price-kline | ✔ |
| [Get Premium Index Price Kline](/market/premium_index_price_kline.go) | GET /v5/market/premium-index-price-kline | ✔ |
| [Get Instruments Info](/market/instruments_info.go) | GET /v5/market/instruments-info | ✔ |
| [Get Orderbook](/market/orderbook.go) | GET /v5/market/orderbook | ✔ |
| [Get Tickers](/market/tickers.go) | GET /v5/market/tickers | ✔ |
| [Get Funding Rate History](/market/.go) | GET /v5/market/funding/history | ✘ |
| [Get Public Recent Trading History](/market/.go) | GET /v5/market/recent-trade | ✘ |
| [Get Open Interest](/market/.go) | GET /v5/market/open-interest | ✘  |
| [Get Historical Volatility](/market/.go) | GET /v5/market/historical-volatility | ✘ |
| [Get Insurance](/market/.go) | GET /v5/market/insurance | ✘ |
| [Get Risk Limit](/market/.go) | GET /v5/market/risk-limit | ✘ |
| [Get Delivery Price](/market/.go) | GET /v5/market/delivery-price | ✘ |
| [Get Long Short Ratio](/market/.go) | GET /v5/market/account-ratio | ✘ |





## Trade

| Endpoint | Description | Status |
|:-------:|:----------- |:------:|
| [Place Order](/trade/place_order.go) | POST /v5/order/create | ✘ |
| [Amend Order](/trade/amend_order.go) | POST /v5/order/amend | ✘ |
| [Cancel Order](/trade/cancel_order.go) | POST /v5/order/cancel | ✘ |
| [Get Open & Closed Orders](/trade/.go) | GET /v5/order/realtime | ✘ |
| [Cancel All Orders](/trade/cancel_all_orders.go) | POST /v5/order/cancel-all | ✘ |
| [Get Order History](/trade/get_order_history.go) | GET /v5/order/history | ✘ |
| [Get Trade History](/trade/.go) | GET /v5/execution/list | ✘ |
| [Batch Place Order](/trade/batch_place_order.go) | POST /v5/order/create-batch | ✘ |
| [Batch Amend Order](/trade/.go) | POST /v5/order/amend-batch | ✘ |
| [Batch Cancel Order](/trade/.go) | POST /v5/order/cancel-batch | ✘ |
| [Get Borrow Quota (Spot)](/trade/.go) | GET /v5/order/spot-borrow-check | ✘ |
| [Set Disconnect Cancel All](/trade/.go) | POST /v5/order/disconnected-cancel-all | ✘ |

## Position

| Endpoint | Description | Status |
|:-------:|:----------- |:------:|
| [Get Position Info](/position/get_position_info.go) | GET /v5/position/list | ✘ |
| [Set Leverage](/position/set_leverage.go) | POST /v5/position/set-leverage | ✘ |
| [Switch Cross/Isolated Margin](/position/.go) | POST /v5/position/switch-isolated | ✘ |
| [Set TP/SL Mode](/position/.go) | POST /v5/position/set-tpsl-mode | ✘ |
| [Switch Position Mode](/position/.go) | POST /v5/position/switch-mode | ✘ |
| [Set Risk Limit](/position/.go) | POST /v5/position/set-risk-limit | ✘ |
| [Set Trading Stop](/position/.go) | POST /v5/position/trading-stop | ✘ |
| [Set Auto Add Margin](/position/.go) | POST /v5/position/set-auto-add-margin | ✘ |
| [Add Or Reduce Margin](/position/.go) | POST /v5/position/add-margin | ✘ |
| [Get Closed PnL](/position/.go) | GET /v5/position/closed-pnl | ✘ |
| [Move Position](/position/.go) | POST /v5/position/move-positions | ✘ |
| [Get Move Position History](/position/.go) | GET /v5/position/move-history | ✘ |
| [Confirm New Risk Limit](/position/.go) | POST /v5/position/confirm-pending-mmr | ✘ |

## Account

| Endpoint | Description | Status |
|:-------:|:----------- |:------:|
| [](/account/.go) | GET /v5/account/ | ✘ |


## WebSocket

| Endpoint | Description | Status |
|:-------:|:----------- |:------:|
| [](/websocket/.go) | GET /v5/ | ✘ |