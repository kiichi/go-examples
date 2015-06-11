package main

import (
	"fmt"
	//"github.com/paulmach/go.geo"
	//"github.com/kiichi/go-examples/gis"
	"net/http"
	//"net/url" // for url.Parse(req.RawQuery)
	//"encodig/json"
)

//https://gobyexample.com/url-parsing
func handler(writer http.ResponseWriter, req *http.Request) {

	fmt.Println("param:%s", req.URL)
	fmt.Fprintf(writer, "{\"test\":\"test\"}")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
