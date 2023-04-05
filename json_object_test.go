package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Married   bool
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Dzaky",
		LastName:  "Mohammad",
		Age:       22,
		Married:   false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}
