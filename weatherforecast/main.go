package main

import (
	"fmt"
	"os"
	m "weatherforecast/meteostation"
)

func main() {
	city := ""
	fmt.Print("select city:")
	_, err := fmt.Fscan(os.Stdin, &city)
	if err != nil {
		fmt.Println("error:", err)
	}
	forecast, err := m.Meteorologist{}.DailyForecast(city, 7)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(forecast.FormatWeatherDaily())

	// forecast, err = m.Meteorologist{}.WeatherForecast(city)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(forecast.FormatWeather())

	// fmt.Println(forecast.GetTemperature())

	fmt.Println(forecast.GetCloudiness())

	// fmt.Println(forecast.GetHumidity())

	// fmt.Println(forecast.GetWind())
}
