package bybit

import (
	"encoding/json"
	"fmt"
	"sort"
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

func ExecuteGet(params string, a any) Response {
	body, _ := Client.GetRequest(params)
	return GetResponse(body, a)
}

/*
func ExecutePost() Response {

}*/

func GetResponse(body []byte, a any) Response {
	var r Response
	err := json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("err:", err)
	}
	Recast(r.Result, &a)
	r.Result = a

	return r
}

func GetParams(a any) string {
	var params string
	b, _ := json.Marshal(a)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {

		x := fmt.Sprint(m[k])

		init_p := k + "=" + x
		if x != "0" && params == "" {
			params = init_p
		}
		if x != "0" && params != init_p {
			params = params + "&" + k + "=" + x
		}

	}

	return params
}
