package meteostation

import (
	"fmt"
)

type WeatherForecast struct {
}

func (wf WeatherForecast) FormatWeather(w Weather) string {
	temp, _, _ := w.GetTemperature()
	speed, gust, direction := w.GetWind()
	sunrise, sunset := w.GetSun()
	var wind string
	if gust != 0 {
		wind = fmt.Sprintf("ветер %v %vм/с с порывами до %vм/с", direction, speed, gust)
	} else {
		wind = fmt.Sprintf("ветер %v %vм/с", direction, speed)
	}
	ws := fmt.Sprintf("Сегодня в городе %v %v, температура воздуха %v°С, %v. Влажность воздуха %v. Восход солнца %v, заход солнца %v",
		w.Name, w.GetCloudiness(), temp, wind, w.GetHumidity(), sunrise, sunset)
	return ws
}
