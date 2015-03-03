package wstest

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	ID        int
}

func Debug(str string) string {
	fmt.Printf("Debug() function called with argument: %s ...\n", str)

	//http://blog.golang.org/json-and-go
	fmt.Printf("==========Serializing Json==============\n")
	req1 := Person{"Kiichi", "Takeuchi", 123}
	data, err := json.Marshal(req1)
	fmt.Printf("%s \nError (%s)\n", data, err)

	fmt.Printf("==========Deserializing Json==============\n")
	var res1 Person
	err2 := json.Unmarshal(data, &res1)
	fmt.Printf("%s \nError (%s)\n", data, err2)

	return str
}
