package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"SamsungS23","price":"2300000"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncodec(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P001",
		"name":  "SamsungS23",
		"price": 2300000,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
