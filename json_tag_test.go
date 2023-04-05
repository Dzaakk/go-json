package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "SamsungS23",
		ImageURL: "http://example.com/samsung23.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"SamsungS23","image_url":"http://example.com/samsung23.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
