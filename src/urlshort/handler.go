package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if dest, found := pathsToUrls[r.URL.Path]; found {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
	//	TODO: Implement this...
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...

	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, fmt.Errorf("fallback handler cannot be nil")
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(data []byte) ([]PathUrl, error) {

	var pathsUrls []PathUrl
	err := yaml.Unmarshal(data, &pathsUrls)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}
	fmt.Printf("Parsed YAML: %+v\n", pathsUrls)
	return pathsUrls, nil
}

func buildMap(pathUrl []PathUrl) map[string]string {
	pathMap := make(map[string]string)
	for _, pu := range pathUrl {
		pathMap[pu.Path] = pu.URL
	}
	if len(pathMap) == 0 {
		fmt.Println("Warning: No paths found in YAML")
	}
	fmt.Printf("Built pathMap: %+v\n", pathMap)
	return pathMap
}

type PathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
