package main

import (
	"fmt"
	m "weatherforecast/meteostation"
)

func main() {
	// NewStation := m.Meteorologist{}.WeatherForecast("Mahilyow")
	NewStation := m.Meteorologist{}.DailyForecast("Mahilyow", 7)
	fmt.Println(NewStation)
	// fmt.Println(NewStation1)
	// fmt.Println(NewStation.FormatWeather())
	// fmt.Println(NewStation1.FormatWeather())
}
