package bybit

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type foo struct {
	Name string
	Code int
}

func verify(a any, fields []string) int {
	code := 0
	b, _ := json.Marshal(a)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	for _, k := range keys {

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

func TestUtils(t *testing.T) {
	f := foo{
		Name: "sss",
	}

	//required fields
	fields := []string{"Name"}

	x := verify(f, fields)
	log.Default().Println(x)

}
