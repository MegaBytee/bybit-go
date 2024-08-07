package bybit

//# version 0.0.3
import (
	"encoding/json"
	"fmt"
	"log"
)

type Response struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     interface{} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

type CallParams struct {
	Method   string
	EndPoint string
	Params   interface{}
	Fields   []string
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

func Call(c *CallParams, a any) {
	MODE := GetEnv("BYBIT_MODE")
	API_KEY := GetEnv("BYBIT_API_KEY")
	API_SECRET := GetEnv("BYBIT_API_SECRET")
	Client := NewConnector().
		SetKeys(API_KEY, API_SECRET).
		SetUrl(MODE).
		SetParams(c.Params, c.Fields).
		SetEndPoint(c.EndPoint)

	var body []byte
	switch c.Method {
	case "GET":
		body, _ = Client.GetRequest()
	case "POST":
		body, _ = Client.PostRequest()
	default:
		panic("method missing...")
	}

	x := getResponse(body, a)

	if x.RetCode != 0 {
		log.Default().Println("code=", x.RetCode)
		log.Default().Println("msg=", x.RetMsg)
	}
}
