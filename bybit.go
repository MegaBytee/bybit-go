package bybit

//# version 0.0.1
import (
	"encoding/json"
	"fmt"
)

const (
	MODE       = "test"
	API_KEY    = ""
	API_SECRET = ""
)

var Client = NewConnector(MODE, API_KEY, API_SECRET)

type Response struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

func Execute(method string, params interface{}, fields []string, a any) Response {
	Client.Params = params
	Client.Required = fields
	var body []byte
	switch method {
	case "GET":
		body, _ = Client.GetRequest()
	case "POST":
		body, _ = Client.PostRequest()
	default:
		panic("method missing...")
	}

	return getResponse(body, a)
}

func getResponse(body []byte, a any) Response {
	var r Response
	err := json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("err:", err)
	}
	Recast(r.Result, &a)
	r.Result = a

	return r
}
