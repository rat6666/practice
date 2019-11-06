package meteostation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const myError = "city not found"
const dailyError = `{"message":"accurate","cod":"200","count":0,"list":[]}`

type Meteorologist struct {
}

func GetBody(city string, days ...string) []byte {
	var cnt string
	prefix := "weather?q="
	// proxyStr := "http://172.26.43.18:3128"
	// proxyURL, err := url.Parse(proxyStr)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	if len(days) != 0 {
		cnt = "&cnt=" + days[0]
		prefix = "find?q="
	}
	urlStr := "http://api.openweathermap.org/data/2.5/" +
		prefix +
		city +
		cnt +
		"&lang=ru" +
		"&units=metric" +
		"&appid=2c19a8c670afc70f2ae7a81f229fce3d"
	// url, err := url.Parse(urlStr)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	// client := &http.Client{Transport: transport}

	// request, err := http.NewRequest("GET", url.String(), nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// response, err := client.Do(request)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer response.Body.Close()
	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	resp, err := http.Get(urlStr)
	if err != nil {
		fmt.Println("Error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error")
	}
	return body
}

func (m Meteorologist) WeatherForecast(city string) (Weather, error) {
	var NewStation Weather
	err := json.Unmarshal(GetBody(city), &NewStation)
	if err != nil {
		fmt.Println(err)
	}
	if NewStation.Message == myError {
		err := errors.New(myError)
		return NewStation, err
	}
	return NewStation, nil
}

func (m Meteorologist) DailyForecast(city string, cnt int) (DailyWeather, error) {
	var NewStation DailyWeather
	body := GetBody(city, strconv.Itoa(cnt))
	err := json.Unmarshal(body, &NewStation)
	if err != nil {
		fmt.Println(err)
	}
	if string(body) == dailyError {
		err := errors.New(myError)
		return NewStation, err
	}
	return NewStation, nil
}
