package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/KananHasanov747/gofiber_weather_app/app/models"
	"github.com/biter777/countries"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

// Safely retrieve the value from .env in right format
func Show[T string | int | bool](key string, _default T) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		return _default
	}

	var result any
	switch any(_default).(type) {
	case string:
		result = val
	case int:
		parsed, err := strconv.Atoi(val)
		if err != nil {
			return _default
		}
		result = parsed
	case bool:
		parsed, err := strconv.ParseBool(val)
		if err != nil {
			return _default
		}
		result = parsed
	default:
		return _default
	}

	return result.(T)
}

func GzipCompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf) // gzipWriter
	_, err := gw.Write(data)
	if err != nil {
		return nil, err
	}
	if err := gw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GzipDecompress(data []byte) (string, error) {
	gr, err := gzip.NewReader(bytes.NewReader(data)) // gzipReader
	if err != nil {
		return "", err
	}
	defer gr.Close()

	var buf bytes.Buffer

	_, err = io.Copy(&buf, gr)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func FetchJSON[T any](endpoint string, paths []string, queries url.Values) (T, error) {
	var result T

	if len(paths) > 0 {
		endpoint = fmt.Sprintf("%s/%s", endpoint, strings.Join(paths, "/"))
	}

	if len(queries) > 0 {
		endpoint = fmt.Sprintf("%s?%s", endpoint, queries.Encode())
	}

	statusCode, resp, err := fasthttp.Get(nil, endpoint)
	if err != nil {
		return result, fmt.Errorf("error making request: %w", err)
	}

	if statusCode != fasthttp.StatusOK {
		return result, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return result, fmt.Errorf("error decoding JSON: %w", err)
	}

	return result, nil
}

func GetLocation(ip string) (map[string]string, error) {
	var result models.Location

	statusCode, resp, err := fasthttp.Get(nil, fmt.Sprintf("https://ipinfo.io/%s/json", ip))

	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if statusCode != fasthttp.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", statusCode)
	}

	err = json.Unmarshal(resp, &result)

	// Check if the result is empty
	if result == (models.Location{}) {
		return nil, fmt.Errorf("fething error")
	}

	loc := strings.Split(result.Loc, ",")
	return map[string]string{
		"city":      result.City,
		"country":   countries.ByName(result.Country).String(),
		"latitude":  loc[0],
		"longitude": loc[1],
	}, nil
}
