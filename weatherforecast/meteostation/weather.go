package meteostation

import (
	"fmt"
	"time"
)

type Weather struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`

	Dt int64

	Main map[string]float64 `json:"main"`

	Sys struct {
		Sunrise int64 `json:"sunrise"`
		Sunset  int64 `json:"sunset"`
	} `json:"sys"`

	Wind map[string]float64 `json:"wind"`

	Name string `json:"name"`
}

func (w Weather) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	return w.Main["temp"], w.Main["temp_min"], w.Main["temp_max"]
}

func (w Weather) GetCloudiness() (description string) {
	return w.Weather[0].Description
}

func (w Weather) GetHumidity() (humidity int) {
	return int(w.Main["humidity"])
}

func (w Weather) GetSun() (sunrise, sunset string) {
	sunrise = time.Unix(w.Sys.Sunrise, 0).Format("15:04")
	sunset = time.Unix(w.Sys.Sunset, 0).Format("15:04")
	return sunrise, sunset
}

func (w Weather) GetDate() (date string) {
	date = time.Unix(w.Dt, 0).Format("2006-01-02")
	return date
}

func (w Weather) GetWind() (speed float64, gust float64, direction string) {
	switch {
	case w.Wind["deg"] >= 341 && w.Wind["deg"] <= 20:
		direction = "северный"
	case w.Wind["deg"] >= 21 && w.Wind["deg"] <= 70:
		direction = "северо-восточный"
	case w.Wind["deg"] >= 71 && w.Wind["deg"] <= 110:
		direction = "восточный"
	case w.Wind["deg"] >= 111 && w.Wind["deg"] <= 160:
		direction = "юго-восточный"
	case w.Wind["deg"] >= 161 && w.Wind["deg"] <= 200:
		direction = "южный"
	case w.Wind["deg"] >= 201 && w.Wind["deg"] <= 250:
		direction = "юго-западный"
	case w.Wind["deg"] >= 251 && w.Wind["deg"] <= 290:
		direction = "западный"
	case w.Wind["deg"] >= 291 && w.Wind["deg"] <= 340:
		direction = "северо-западный"
	}
	return w.Wind["speed"], w.Wind["gust"], direction
}

func (w Weather) FormatWeather() string {
	temp, _, _ := w.GetTemperature()
	speed, gust, direction := w.GetWind()
	sunrise, sunset := w.GetSun()
	var wind string
	if gust != 0 {
		wind = fmt.Sprintf("ветер %v %vм/с с порывами до %vм/с", direction, speed, gust)
	} else {
		wind = fmt.Sprintf("ветер %v %vм/с", direction, speed)
	}
	ws := fmt.Sprintf("Сегодня, %v в городе %v %v, температура воздуха %v°С, %v."+
		"Влажность воздуха %v%v. Восход солнца %v, заход солнца %v",
		w.GetDate(), w.Name, w.GetCloudiness(), temp, wind, w.GetHumidity(),
		"%", sunrise, sunset)
	return ws
}
