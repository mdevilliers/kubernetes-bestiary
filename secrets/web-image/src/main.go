package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>")
	listAllEnvironmentalVars(w)
	loadSecret(w, "username")
	loadSecret(w, "password")
	loadSecret(w, "another-secret")
	fmt.Fprint(w, "</body><html>")
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func loadSecret(w http.ResponseWriter, key string) {

	path := "/etc/foo/"

	secret, err := ioutil.ReadFile(path + key)

	if err != nil {
		fmt.Fprintf(w, "error reading %s : %s", key, err.Error())
		return
	}

	fmt.Fprintf(w, "<br/>ssh!, Key : %s, Secret : %s", key, secret)

}

func listAllEnvironmentalVars(w http.ResponseWriter) {
	for _, e := range os.Environ() {
		fmt.Fprintf(w, "%s<br/>", e)
	}
}
