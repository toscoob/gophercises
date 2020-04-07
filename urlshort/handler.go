package urlshort

import (
	//"fmt"
	"net/http"
	"gopkg.in/yaml.v2"
	"encoding/json"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	res := func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
	return res
}

type urlentry struct {
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}

func parseYAML(yml []byte) ([]urlentry, error){
	var u []urlentry
	//fmt.Println(yml)
	err := yaml.Unmarshal(yml, &u)
	//fmt.Println(u)
	return u, err
}

func buildMap(parsedYml []urlentry) (map[string]string, error){
	//m := make([]urlentry, len(parsedYml))
	var m map[string]string
	m = make(map[string]string)
	for _, v := range parsedYml {
		m[v.Path] = v.Url
	}
	return m, nil
}
// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
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
		return nil, err
	}
	pathMap, _ := buildMap(parsedYaml)
	//fmt.Println(pathMap)
	return MapHandler(pathMap, fallback), nil

	//return fallback.ServeHTTP, nil
}

func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var u []urlentry
	err := json.Unmarshal(jsn, &u)
	if err != nil {
		return nil, err
	}

	pathMap, _ := buildMap(u)
	//fmt.Println(pathMap)
	return MapHandler(pathMap, fallback), nil

	//return fallback.ServeHTTP, nil
}
