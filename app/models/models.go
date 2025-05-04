package models

type Search struct {
	Name       string  `json:"name"`
	Country    string  `json:"country"`
	Latitude   float32 `json:"latitude"`
	Longitude  float32 `json:"longitude"`
	Population uint32  `json:"population"`
}

type SearchResponse struct {
	Results []Search `json:"results"`
}

type Hourly struct {
	Apparent_temperature      []float32 `json:"apparent_temperature"`
	Precipitation_probability []int     `json:"precipitation_probability"`
	Rain                      []float32 `json:"rain"`
	Relative_humidity_2m      []int     `json:"relative_humidity_2m"`
	Temperature_2m            []float32 `json:"temperature_2m"`
	Time                      []string  `json:"time"`
	Weather_code              []int     `json:"weather_code"`
	Wind_speed_10m            []float32 `json:"wind_speed_10m"`
}

type Daily struct {
	Temperature_2m_max []float32 `json:"temperature_2m_max"`
	Temperature_2m_min []float32 `json:"temperature_2m_min"`
	Time               []string  `json:"time"`
	Uv_index_max       []float32 `json:"uv_index_max"`
	Weather_code       []int     `json:"weather_code"`
}

type WeatherResponse struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Hourly    Hourly  `json:"hourly"`
	Daily     Daily   `json:"daily"`
	Timezone  string  `json:"timezone"`
}
