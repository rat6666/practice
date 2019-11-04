package meteostation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Meteorologist struct {
}

func (m Meteorologist) WeatherForecast(city string) Weather {
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
