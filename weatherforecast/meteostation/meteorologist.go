package meteostation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Meteorologist struct {
	cnt int
}

func (Meteorologist) WeatherForecast(city string) Weather {
	url := "http://api.openweathermap.org/data/2.5/weather?q=" +
		city +
		"&lang=ru" +
		"&units=metric" +
		"&appid=2c19a8c670afc70f2ae7a81f229fce3d"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error")
	}
	var NewStation Weather
	err = json.Unmarshal(body, &NewStation)
	if err != nil {
		fmt.Println("error:", err)
	}
	return NewStation
}

func (Meteorologist) DailyForecast(city string, cnt int) Weather {
	url := "http://api.openweathermap.org/data/2.5/find?q=" +
		city +
		"&units=metric" +
		"&lang=ru" +
		"&cnt=" +
		strconv.Itoa(cnt) +
		"&appid=2c19a8c670afc70f2ae7a81f229fce3d"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error")
	}
	type test interface{}
	var itest test
	err = json.Unmarshal(body, &itest)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(itest)

	var NewStation Weather
	err = json.Unmarshal(body, &NewStation)
	if err != nil {
		fmt.Println("error:", err)
	}
	return NewStation
}
