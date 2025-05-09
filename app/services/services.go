package services

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/KananHasanov747/gofiber_weather_app/app/models"
	"github.com/KananHasanov747/gofiber_weather_app/cache/redis"
	"github.com/KananHasanov747/gofiber_weather_app/openmeteo"
	"github.com/KananHasanov747/gofiber_weather_app/pkg/utils"
	"github.com/goccy/go-json"
)

const (
	weatherURL = "https://api.open-meteo.com/v1/forecast"
	geoURL     = "https://geocoding-api.open-meteo.com/v1/search"
)

func GetSearchData(city string) ([]byte, error) {
	city = strings.ToLower(city)
	cache_key := "search_" + strings.Replace(city, "%20", "_", -1)

	raw, err := redis.Store.Get(cache_key)
	if err == nil && len(raw) > 0 {
		return raw, nil
	}

	params := url.Values{
		"name":     {strings.Replace(city, "%20", " ", -1)},
		"count":    {"10"},
		"language": {"en"},
		"format":   {"json"},
	}

	list, err := utils.FetchJSON[models.SearchResponse](geoURL, nil, params)
	if err != nil {
		return nil, fmt.Errorf("GetSearchData FetchJSON error: %v\n", err)
	}

	listBytes, err := json.Marshal(list.Results)
	if err != nil {
		return nil, fmt.Errorf("GetSearchData marshal error: %w", err)
	}
	compressed, err := utils.GzipCompress(listBytes)
	if err == nil && len(compressed) > 0 {
		now := time.Now()
		remaining := 3600 - (now.Minute()*60 + now.Second()) // the remaining time left to the next hour
		redis.Store.Set(cache_key, compressed, time.Duration(remaining)*time.Second)
	}

	return compressed, nil
}

func GetWeatherData(city, country, lat, lon string) ([]byte, error) {
	cache_key := fmt.Sprintf("weather_%s_%s", lat, lon)

	raw, err := redis.Store.Get(cache_key)
	if err == nil && len(raw) > 0 {
		return raw, nil
	}

	daily := []string{
		"weather_code",
		"temperature_2m_max",
		"temperature_2m_min",
		"uv_index_max",
	}

	hourly := []string{
		"temperature_2m",
		"relative_humidity_2m",
		"apparent_temperature",
		"rain",
		"precipitation_probability",
		"wind_speed_10m",
		"weather_code",
		"is_day",
	}

	params := url.Values{
		"latitude":       {lat},
		"longitude":      {lon},
		"daily":          {strings.Join(daily, ",")},
		"hourly":         {strings.Join(hourly, ",")},
		"timezone":       {"auto"},
		"forecast_hours": {"24"},
	}

	data, err := utils.FetchJSON[models.WeatherResponse](weatherURL, nil, params)
	if err != nil {
		return nil, fmt.Errorf("GetWeatherData FetchJSON error: %v\n", err)
	}

	for i, w_code := range data.Hourly.Weather_code {
		icon := openmeteo.Icons(w_code, data.Hourly.Is_day[i])
		data.Hourly.Icon_url = append(data.Hourly.Icon_url, "/static/assets/icons/"+icon+".svg")
		data.Hourly.Description = append(data.Hourly.Description, icon)
	}

	for _, w_code := range data.Daily.Weather_code {
		icon := openmeteo.Icons(w_code, 1)
		data.Daily.Icon_url = append(data.Daily.Icon_url, "/static/assets/icons/"+icon+".svg")
		data.Daily.Description = append(data.Daily.Description, icon)
	}

	data.City, data.Country = city, country
	data.Current = models.Current{
		Temperature_2m:       data.Hourly.Temperature_2m[0],
		Apparent_temperature: data.Hourly.Apparent_temperature[0],
		Icon_url:             data.Hourly.Icon_url[0],
		Description:          data.Hourly.Description[0],
		Rain:                 data.Hourly.Rain[0],
		Wind_speed_10m:       data.Hourly.Wind_speed_10m[0],
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("GetWeatherData marshal error: %w", err)
	}

	compressed, err := utils.GzipCompress(dataBytes)
	if err == nil && len(compressed) > 0 {
		now := time.Now()
		remaining := 3600 - (now.Minute()*60 + now.Second()) // the remaining time left to the next hour
		redis.Store.Set(cache_key, compressed, time.Duration(remaining)*time.Second)
	}

	return compressed, nil
}
