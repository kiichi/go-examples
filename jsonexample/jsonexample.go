package jsonexample

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	ID        int
}

type PhoneBook map[string]string

func BasicJsonExample() {
	fmt.Printf("\n\n[BasicJsonExample]\n")

	//http://blog.golang.org/json-and-go
	fmt.Printf("==========Serializing Json==============\n")
	req1 := Person{"Kiichi", "Takeuchi", 123}
	data, err := json.Marshal(req1)
	if err != nil {
		fmt.Printf("%s \n", err)
	}
	fmt.Printf("%s \n", data)

	fmt.Printf("==========Deserializing Json==============\n")
	var res1 Person
	err2 := json.Unmarshal(data, &res1)
	if err2 != nil {
		fmt.Printf("error %s", err2)
	}

	fmt.Printf("%s \n", res1.FirstName)
}

func MapToJsonExample() {
	fmt.Printf("\n\n[MapToJsonExample]\n")

	fmt.Printf("==========Serializing Json==============\n")
	phones := make(map[string]string)
	phones2 := make(map[string]string)
	phones["Kiichi"] = "123-333-0998"
	phones["John"] = "998-333-2288"
	phones["Mary"] = "39028-38297"
	data, err := json.Marshal(phones)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s \n", data)

	fmt.Printf("==========Deserializing Json==============\n")
	err2 := json.Unmarshal(data, &phones2)
	if err2 != nil {
		fmt.Printf("%s \n", err2)
	}
	fmt.Printf("%s \n", phones2)
}
