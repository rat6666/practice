package meteostation

import "time"

type Weather struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`

	Main map[string]float64 `json:"main"`

	Sys struct {
		Sunrise int64 `json:"sunrise"`
		Sunset  int64 `json:"sunset"`
	} `json:"sys"`

	Wind map[string]int `json:"wind"`

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

func (w Weather) GetWind() (speed int, gust int, direction string) {
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
