package models

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Loc     string `josn:"loc"`
}

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

type Current struct {
	Temperature_2m       float32 `json:"temperature_2m"`
	Apparent_temperature float32 `json:"apparent_temperature"`
	Icon_url             string  `json:"icon_url"`
	Description          string  `json:"description"`
	Rain                 float32 `json:"rain"`
	Wind_speed_10m       float32 `josn:"wind_speed_10m"`
}

type Hourly struct {
	Apparent_temperature      []float32 `json:"apparent_temperature"`
	Is_day                    []int     `json:"is_day"`
	Icon_url                  []string  `json:"icon_url"`
	Description               []string  `json:"description"`
	Precipitation_probability []int     `json:"precipitation_probability"`
	Temperature_2m            []float32 `json:"temperature_2m"`
	Relative_humidity_2m      []int     `json:"relative_humidity_2m"`
	Rain                      []float32 `json:"rain"`
	Time                      []string  `json:"time"`
	Weather_code              []int     `json:"weather_code"`
	Wind_speed_10m            []float32 `json:"wind_speed_10m"`
}

type Daily struct {
	Temperature_2m_max []float32 `json:"temperature_2m_max"`
	Temperature_2m_min []float32 `json:"temperature_2m_min"`
	Day_of_week        []string  `json:"day_of_week"`
	Time               []string  `json:"time"`
	Icon_url           []string  `json:"icon_url"`
	Description        []string  `json:"description"`
	Uv_index_max       []float32 `json:"uv_index_max"`
	Weather_code       []int     `json:"weather_code"`
}

type WeatherResponse struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Current   Current `json:"current"`
	Hourly    Hourly  `json:"hourly"`
	Daily     Daily   `json:"daily"`
}
