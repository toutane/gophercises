package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

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
	parsedYaml, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}

	buildedMap, err := buildMap(parsedYaml)
	if err != nil {
		return nil, err
	}

	return MapHandler(*buildedMap, fallback), nil
}

type Item struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func parseYaml(yml []byte) (*[]Item, error) {
	var result []Item

	err := yaml.Unmarshal(yml, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func buildMap(parsedYaml *[]Item) (*map[string]string, error) {
	result := make(map[string]string)
	for _, item := range *parsedYaml {
		result[item.Path] = item.Url
	}

	return &result, nil
}
