package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/toutane/gophercises/urlshort/urlshort"
)

func main() {
	mux := defaultMux()

	//Build the MalHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-gpdoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-sodoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//Build the YAMLHandler using the mapHandler as the fallback.
	filename := "data.yaml"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := urlshort.YAMLHandler(bytes, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8000")
	http.ListenAndServe(":8000", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
