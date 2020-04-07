package main

import (
	"flag"
	"fmt"
	"github.com/gophercises/urlshort"
	"io/ioutil"
	"net/http"
)

// go run main/main.go

func main() {
	yamlFilename := flag.String("y", "", "a yaml file in the format of path: 'path', url: 'url'")
	jsonFilename := flag.String("j", "", "a yaml file in the format of path: 'path', url: 'url'")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	var handler http.HandlerFunc

	if *jsonFilename != "" {
		content, err := ioutil.ReadFile(*jsonFilename)
		if err != nil {
			panic(err)
		}
		jsn := string(content)
		handler, err = urlshort.JSONHandler([]byte(jsn), mapHandler)
		if err != nil {
			panic(err)
		}
	} else if *yamlFilename != "" {
		content, err := ioutil.ReadFile(*yamlFilename)
		if err != nil {
			panic(err)
		}
		yaml := string(content)
		handler, err = urlshort.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			panic(err)
		}
	} else {
		var err error
		yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
		//repeated code, but ok
		handler, err = urlshort.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			panic(err)
		}
	}
/*
	fmt.Println(yaml)
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
*/
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
	//http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
