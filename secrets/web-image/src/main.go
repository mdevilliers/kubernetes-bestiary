package main

import (
	"encoding/base64"
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

	encoded, err := ioutil.ReadFile(path + key)

	if err != nil {
		fmt.Fprintf(w, "error reading %s : %s", key, err.Error())
		return
	}
	unencoded, err := base64.StdEncoding.DecodeString(string(encoded))

	if err != nil {
		fmt.Fprintf(w, "error decoding %s : %s", string(encoded), err.Error())
		return
	}

	fmt.Fprintf(w, "<br/>ssh!, Key : %s, Secret : %s from : %s", key, unencoded, string(encoded))

}

func listAllEnvironmentalVars(w http.ResponseWriter) {
	for _, e := range os.Environ() {
		fmt.Fprintf(w, "%s<br/>", e)
	}
}
