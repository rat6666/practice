package meteostation

import (
	"encoding/json"
	"errors"
	"fmt"
)

const myError = "city not found"

type Meteorologist struct {
}

func (Meteorologist) WeatherForecast(city string) (Weather, error) {
	body := GetUrlCurrent(city)
	var NewStation Weather
	err := json.Unmarshal(body, &NewStation)
	if err != nil {
		fmt.Println(err)
	}
	if NewStation.Message == myError {
		err := errors.New(myError)
		return NewStation, err
	}
	return NewStation, nil
}

func (Meteorologist) DailyForecast(city string, cnt int) (Weather, error) {
	body := GetUrlDaily(city, cnt)
	var NewStation Weather
	err := json.Unmarshal(body, &NewStation)
	if err != nil {
		fmt.Println(err)
	}
	if NewStation.Message == myError {
		err := errors.New(myError)
		return NewStation, err
	}
	return NewStation, nil
}
