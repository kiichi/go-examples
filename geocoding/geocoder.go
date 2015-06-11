package geocoding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ReadSchools() []interface{} {
	var data map[string]interface{}
	content, err := ioutil.ReadFile("./schools.json")
	if err != nil {
		fmt.Println("error reading file : ", err)
	}
	//fmt.Println("%s", file)

	err2 := json.Unmarshal(content, &data)

	if err2 != nil {
		fmt.Println("error json parse : ", err)
	}

	//		fmt.Println("data", data["results"].(map[string]interface{})["Schools"].([]interface{})[0].(map[string]interface{})["STREET"])
	return data["results"].(map[string]interface{})["Schools"].([]interface{})
}

func Geocode(schools []interface{}) {
	for i := 0; i < len(schools); i++ {
		dict := schools[i].(map[string]interface{})
		fmt.Printf("%s,%s\n", dict["STREET"], dict["CITY"])
		url := fmt.Sprint("http://open.mapquestapi.com/nominatim/v1/search.php?format=json&q=", url.QueryEscape(dict["STREET"].(string)), ",", url.QueryEscape(dict["CITY"].(string)), url.QueryEscape(" NY USA"))
		fmt.Println(url)
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		//	fmt.Println(body)

		var data []interface{}
		err2 := json.Unmarshal(body, &data)
		if err2 != nil {
			fmt.Println(err2)
		}
		if len(data) == 0 {
			fmt.Println("no data")
		} else {
			coord := data[i].(map[string]interface{})
			fmt.Println(coord["lat"], coord["lon"])
			schools[i].(map[string]interface{})["lat"] = coord["lat"]
			schools[i].(map[string]interface{})["lon"] = coord["lon"]
			fmt.Println(schools[i])
		}
	}
	content, err := json.Marshal(schools)
	if err != nil {
		fmt.Println("marshal error ", err)
	}
	err2 := ioutil.WriteFile("./schools_geocoded.json", content, 0644)
	if err2 != nil {
		fmt.Println(err2)
	}
}
