package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Item struct {
	Path string
	Url  string
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		requestUrl := r.URL.Path

		if pathsToUrls[requestUrl] != "" {
			url := pathsToUrls[requestUrl]
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(f)
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseData("yaml", yml)
	if err != nil {
		return nil, err
	}

	buildedMap, err := buildMap(parsedYaml)
	if err != nil {
		return nil, err
	}

	return MapHandler(*buildedMap, fallback), nil
}

func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJson, err := parseData("json", jsn)
	if err != nil {
		return nil, err
	}

	buildedMap, err := buildMap(parsedJson)
	if err != nil {
		return nil, err
	}

	return MapHandler(*buildedMap, fallback), nil
}

func parseData(ext string, data []byte) (*[]Item, error) {
	var result []Item

	if ext == "yaml" {
		err := yaml.Unmarshal(data, &result)
		if err != nil {
			return nil, err
		}
	} else {
		err := json.Unmarshal(data, &result)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func buildMap(parsedData *[]Item) (*map[string]string, error) {
	result := make(map[string]string)

	for _, item := range *parsedData {
		result[item.Path] = item.Url
	}

	return &result, nil
}
