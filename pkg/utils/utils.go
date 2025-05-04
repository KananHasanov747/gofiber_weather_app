package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/goccy/go-json"
)

func Convert[T any](val string, result *T) error {
	switch any(result).(type) {
	case *string:
		*result = any(val).(T)
		return nil
	case *int:
		// Conversion to int (example)
		var temp int
		_, err := fmt.Sscanf(val, "%d", &temp)
		if err == nil {
			*result = any(temp).(T)
		}
		return err
	default:
		return fmt.Errorf("unsupported type: %T", result)
	}
}

func Show[T any](key string, def T) T {
	val, ok := os.LookupEnv(key)
	if ok {
		var result T
		if err := Convert(val, &result); err == nil {
			return result
		}
	}
	return def
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
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

	resp, err := http.Get(endpoint)
	if err != nil {
		return result, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("error decoding JSON: %w", err)
	}

	return result, nil
}
