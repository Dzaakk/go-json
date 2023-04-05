package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "Dzaky",
		LastName:  "Mohammad",
		Hobbies:   []string{"Archery", "Coding", "Billiard"},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonStirng := `{"FirstName":"Dzaky","LastName":"Mohammad","Age":0,"Married":false,"Hobbies":["Archery","Coding","Billiard"]}`
	jsonBytes := []byte(jsonStirng)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Dzaky",
		Addresses: []Address{
			{
				Street:     "Jalan 1",
				Country:    "Indonesia",
				PostalCode: "123",
			},
			{
				Street:     "Jalan 2",
				Country:    "Indonesia",
				PostalCode: "456",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Dzaky","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan 1","Country":"Indonesia","PostalCode":"123"},{"Street":"Jalan 2","Country":"Indonesia","PostalCode":"456"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)

}
func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan 1","Country":"Indonesia","PostalCode":"123"},{"Street":"Jalan 2","Country":"Indonesia","PostalCode":"456"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
func TestOnlyJsonArrayComplex(t *testing.T) {
	Addresses := []Address{
		{
			Street:     "Jalan 1",
			Country:    "Indonesia",
			PostalCode: "123",
		},
		{
			Street:     "Jalan 2",
			Country:    "Indonesia",
			PostalCode: "456",
		},
	}

	bytes, _ := json.Marshal(Addresses)
	fmt.Println(string(bytes))
}
