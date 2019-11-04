package main

import (
	"fmt"
	m "weatherforecast/meteostation"
)

func main() {
	NewStation := m.Meteorologist{}.WeatherForecast("Mahilyow")
	x := m.WeatherForecast{}
	fmt.Println(x.FormatWeather(NewStation))
}
