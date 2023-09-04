package bybit

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"sort"
	"strconv"
	"time"
)

const (
	TESTNET_URL = "https://api-testnet.bybit.com"
	MAINNET_URL = "https://api.bybit.com"
)

type Connector struct {
	Params     interface{} //request params
	Required   []string    //required fields
	Url        string      //main, or test api url
	EndPoint   string      //api endpoint method
	ApiKey     string      //user api key
	ApiSecret  string      //user api secret
	RecvWindow string      //recvwindow= 5000 default
	Signature  string      //user signature
	httpClient *http.Client
}

func NewConnector() *Connector {

	return &Connector{
		RecvWindow: "5000",
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}

}

func (c *Connector) SetKeys(key, secret string) *Connector {
	c.ApiKey = key
	c.ApiSecret = secret
	return c
}

func (c *Connector) SetEndPoint(value string) *Connector {
	c.EndPoint = value
	return c
}
func (c *Connector) SetUrl(mode string) *Connector {
	switch mode {
	case "main":
		c.Url = MAINNET_URL
	case "test":
		c.Url = TESTNET_URL
	}

	return c
}

func (c *Connector) SetParams(p interface{}, required []string) *Connector {
	c.Params = p
	c.Required = required
	return c

}

func getParams(c *Connector) string {

	var params string
	b, _ := json.Marshal(c.Params)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)

	if verifyRequired(m, c.Required) != 0 {
		log.Default().Println("params not valid")
		return ""
	}

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

func verifyRequired(m map[string]interface{}, fields []string) int {
	code := 0

	for k := range m {
		for _, x := range fields {
			if k == x {

				y := fmt.Sprint(m[k])
				if y == "" || y == "0" {
					code++
				}

			}
		}

	}

	return code
}

func (c *Connector) GetRequest() ([]byte, int) {

	params := getParams(c)

	now := time.Now()

	unixNano := now.UnixNano()
	time_stamp := unixNano / 1000000
	hmac256 := hmac.New(sha256.New, []byte(c.ApiSecret))
	hmac256.Write([]byte(strconv.FormatInt(time_stamp, 10) + c.ApiKey + c.RecvWindow + params))
	c.Signature = hex.EncodeToString(hmac256.Sum(nil))
	request, error := http.NewRequest(http.MethodGet, c.Url+c.EndPoint+"?"+params, nil)
	if error != nil {
		panic(error)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-BAPI-API-KEY", c.ApiKey)
	request.Header.Set("X-BAPI-SIGN", c.Signature)
	request.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(time_stamp, 10))
	request.Header.Set("X-BAPI-SIGN-TYPE", "2")
	request.Header.Set("X-BAPI-RECV-WINDOW", c.RecvWindow)
	reqDump, err := httputil.DumpRequestOut(request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Request Dump:\n%s", string(reqDump))
	response, error := c.httpClient.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	elapsed := time.Since(now).Seconds()
	fmt.Printf("\n%s took %v seconds \n", c.EndPoint, elapsed)
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	return body, response.StatusCode
}

func (c *Connector) PostRequest() ([]byte, int) {

	params := c.Params
	now := time.Now()
	unixNano := now.UnixNano()
	time_stamp := unixNano / 1000000
	jsonData, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
	}
	hmac256 := hmac.New(sha256.New, []byte(c.ApiSecret))
	hmac256.Write([]byte(strconv.FormatInt(time_stamp, 10) + c.ApiKey + c.RecvWindow + string(jsonData[:])))
	c.Signature = hex.EncodeToString(hmac256.Sum(nil))
	request, error := http.NewRequest(http.MethodPost, c.Url+c.EndPoint, bytes.NewBuffer(jsonData))
	if error != nil {
		panic(error)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-BAPI-API-KEY", c.ApiKey)
	request.Header.Set("X-BAPI-SIGN", c.Signature)
	request.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(time_stamp, 10))
	request.Header.Set("X-BAPI-SIGN-TYPE", "2")
	request.Header.Set("X-BAPI-RECV-WINDOW", c.RecvWindow)
	reqDump, err := httputil.DumpRequestOut(request, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Request Dump:\n%s", string(reqDump))
	response, error := c.httpClient.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	elapsed := time.Since(now).Seconds()
	fmt.Printf("\n%s took %v seconds \n", c.EndPoint, elapsed)
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return body, response.StatusCode
}
