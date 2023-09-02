package bybit

import (
	"encoding/json"
	"strconv"
)

// recast a to b
func Recast(a, b any) error {
	js, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return json.Unmarshal(js, b)
}

func Serialize(s any) []byte {
	data, err := json.Marshal(s)
	if err != nil {
		print(err)
		return []byte(nil)
	}

	return data
}

func DeSerialize(data []byte) any {
	var a any
	err := json.Unmarshal(data, &a)
	if err != nil {
		print(err)
	}

	return a

}

func StringToInt64(value string) int64 {
	p, _ := strconv.ParseInt(value, 10, 64)
	return p
}
func StringToInt32(value string) int32 {
	p, _ := strconv.ParseInt(value, 10, 32)
	return int32(p)
}
func StringToFloat64(value string) float64 {
	p, _ := strconv.ParseFloat(value, 64)
	return p
}

func StringToBool(value string) bool {
	p, _ := strconv.ParseBool(value)
	return p
}
