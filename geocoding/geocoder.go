package geocoding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
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

// simple version for example
func Geocode(schools []interface{}) []interface{} {
	//for i := 0; i < len(schools); i++ {
	// just do first one for test.
	for i := 0; i < 1; i++ {
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
	return schools
}
func WriteSchools(schools []interface{}) {
	content, err := json.Marshal(schools)
	if err != nil {
		fmt.Println("marshal error ", err)
	}
	err2 := ioutil.WriteFile("./schools_geocoded.json", content, 0644)
	if err2 != nil {
		fmt.Println(err2)
	}
}

func WriteSchoolsSQL(schools []interface{}) {
	str := ""
	for i := 0; i < len(schools); i++ {
		var row = schools[i].(map[string]interface{})
		if row["lat"] != nil {
			str = fmt.Sprint(str, "UPDATE Ambassador SET lat='", row["lat"], "', lng='", row["lng"], "' WHERE ENTITY_CD='", row["ENTITY_CD"], "';\n")
		}
	}
	err2 := ioutil.WriteFile("./schools_geocoded.sql", []byte(str), 0644)
	if err2 != nil {
		fmt.Println(err2)
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Resume version
// check if there is a schools_geocoded.json file or not
// if it's there, read it and continue from where it left off
func GeocodeR(schools []interface{}) []interface{} {

	//var cache []interface{}
	cache := ReadSchoolsR()
	if cache != nil {
		schools = cache
	}

	// just do first one for test.
	//for i := 0; i < 3; i++ {
	for i := 0; i < len(schools); i++ {
		// skip if it's been geocoded already
		tmp := schools[i].(map[string]interface{})
		// reprocess unmatched record or not
		//if tmp["lat"] != nil || (tmp["match"] != nil && tmp["match"] == false) {
		if tmp["lat"] != nil {
			continue
		}
		dict := schools[i].(map[string]interface{})
		var street = strings.TrimRight(dict["STREET"].(string), "-")
		fmt.Printf("%s,%s\n", street, dict["REGION"])
		url := fmt.Sprint("http://maps.googleapis.com/maps/api/geocode/json?address=", url.QueryEscape(street), ",", url.QueryEscape(dict["REGION"].(string)), url.QueryEscape(" NY USA"), "&sensor=false")
		fmt.Println(url)
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		//	fmt.Println(body)

		var data map[string]interface{}
		err2 := json.Unmarshal(body, &data)
		if err2 != nil {
			fmt.Println(err2)
		}
		if data["status"] != "OK" {
			fmt.Println("no data")
			schools[i].(map[string]interface{})["match"] = false
		} else {
			coord := data["results"].([]interface{})[0].(map[string]interface{})["geometry"].(map[string]interface{})["location"].(map[string]interface{})
			fmt.Println(coord["lat"], coord["lon"])
			schools[i].(map[string]interface{})["lat"] = coord["lat"]
			schools[i].(map[string]interface{})["lng"] = coord["lng"]
			schools[i].(map[string]interface{})["match"] = true
			fmt.Println(i, "/", len(schools), " ", schools[i])
		}

		// Always write the latest schools data
		WriteSchools(schools)
		time.Sleep(500 * time.Millisecond)
	}
	return schools
}

func ReadSchoolsR() []interface{} {
	var data []interface{}
	content, err := ioutil.ReadFile("./schools_geocoded.json")
	if err != nil {
		fmt.Println("nothing to read. error reading file : ", err)
		return nil
	}

	err2 := json.Unmarshal(content, &data)

	if err2 != nil {
		fmt.Println("error json parse : ", err)
	}

	return data
}
