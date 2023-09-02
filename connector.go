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
	"strconv"
	"time"
)

const (
	TESTNET_URL = "https://api-testnet.bybit.com"
	MAINNET_URL = "https://api.bybit.com"
)

type Connector struct {
	Url        string
	EndPoint   string
	ApiKey     string
	ApiSecret  string
	RecvWindow string
	Signature  string
}

var http_client *http.Client = &http.Client{Timeout: 10 * time.Second}

func NewConnector(mode, api_key, api_secret string) Connector {

	c := Connector{
		ApiKey:     api_key,
		ApiSecret:  api_secret,
		RecvWindow: "5000",
	}
	c.SetUrl(mode)
	return c
}

func (c *Connector) SetEndPoint(value string) {
	c.EndPoint = value
}
func (c *Connector) SetUrl(mode string) {
	switch mode {
	case "main":
		c.Url = MAINNET_URL
	case "test":
		c.Url = TESTNET_URL
	}
}
func (c *Connector) GetRequest(params string) ([]byte, int) {

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
	response, error := http_client.Do(request)
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

func (c *Connector) PostRequest(params interface{}) ([]byte, int) {

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
	response, error := http_client.Do(request)
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
