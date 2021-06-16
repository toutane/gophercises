package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/toutane/gophercises/urlshort/urlshort"
)

func main() {
	mux := defaultMux()

	filename := flag.String("file", "data.yaml", "File where to get the data.")
	flag.Parse()

	//First part of exercise.
	//Build the MalHandler using the mux as the fallback
	/*
		pathsToUrls := map[string]string{
			"/urlshort-gpdoc": "https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-sodoc":     "https://godoc.org/gopkg.in/yaml.v2",
		}

		mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	*/

	//Second part of exercise.
	//Build the YAMLHandler using the mapHandler as the fallback.
	bytes, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}

	//Get the file extension.
	arr := strings.Split(*filename, ".")
	ext := arr[len(arr)-1]

	var handler http.Handler
	if ext == "yaml" {
		yamlHandler, err := urlshort.YAMLHandler(bytes, mux)
		if err != nil {
			panic(err)
		}
		handler = yamlHandler
	} else {
		jsonHandler, err := urlshort.JSONHandler(bytes, mux)
		if err != nil {
			panic(err)
		}
		handler = jsonHandler
	}

	fmt.Println("Starting the server on :8000")
	http.ListenAndServe(":8000", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
