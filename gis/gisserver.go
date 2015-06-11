package main

import (
	"encoding/json"
	"fmt"
	"github.com/paulmach/go.geo"
	"net/http"
)

type Res struct {
	Status  string
	Result  bool
	Message string
}

//https://gobyexample.com/url-parsing
func handler(writer http.ResponseWriter, req *http.Request) {
	payload := req.FormValue("payload")
	fmt.Printf("payload via x-www-form-urlencoded: %s\n", payload)

	line1 := geo.NewLine(geo.NewPoint(-73.942799, 40.7459951), geo.NewPoint(-73.937349, 40.748783))
	line2 := geo.NewLine(geo.NewPoint(-73.944494, 40.751250), geo.NewPoint(-73.937010, 40.743877))

	//b, _ := json.Marshal(geo.NewPoint(-73.942799, 40.7459951))
	var result bool = line1.Intersects(line2)
	res := Res{"ok", result, payload}
	data, _ := json.Marshal(res)
	fmt.Fprintf(writer, "%s", data)
}

func main() {
	http.HandleFunc("/intersect/", handler)
	http.ListenAndServe(":8080", nil)
}
