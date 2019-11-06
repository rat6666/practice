package main

import (
	"fmt"
	"os"
	m "weatherforecast/meteostation"
)

func main() {
	city := ""
	fmt.Print("Input city:")
	_, err := fmt.Fscan(os.Stdin, &city)
	if err != nil {
		fmt.Println("Error:", err)
	}

	forecast, err := m.Meteorologist{}.WeatherForecast(city)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
	dailyforecast, err := m.Meteorologist{}.DailyForecast(city, 7)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}

	fmt.Println(forecast.FormatWeather())

	fmt.Println(dailyforecast.FormatDailyWeather())

	fmt.Println(forecast.GetTemperature())

	fmt.Println(forecast.GetCloudiness())

	fmt.Println(forecast.GetHumidity())

	fmt.Println(forecast.GetWind())
}
