package main

import (
	"fmt"
	m "weatherforecast/meteostation"
)

func main() {
	NewStation := m.Meteorologist{}.WeatherForecast("Tokio")
	fmt.Println(NewStation.FormatWeather())
}
