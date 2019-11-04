package meteostation

type Station struct {
	Weather Weather            `json:"weather"`
	Main    map[string]float64 `json:"main"`
	Sys     Sys                `json:"sys"`
	Wind    map[string]float64 `json:"wind"`
	Name    string             `json:"name"`
}

type Weather []struct {
	Description string `json:"description"`
}

type Sys struct {
	Sunrise int64 `json:"sunrise"`
	Sunset  int64 `json:"sunset"`
}

func (s Station) GetTemperature() (temp float64, tempMin float64, tempMax float64) {
	return s.Main["temp"], s.Main["temp_min"], s.Main["temp_max"]
}

func (s Station) GetCloudiness() (description string) {
	return s.Weather[0].Description
}

func (s Station) GetHumidity() (humidity int) {
	return int(s.Main["humidity"])
}

// fmt.Println((time.Unix(NewStation.Sys.Sunrise, 0)).Format("15:55")
