package main

import (
	"fmt"
	rediscli "github.com/garyburd/redigo/redis"
	"log"
	"net/http"
)

func main() {

	conn, err := rediscli.Dial("tcp", "localhost:6379")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer conn.Close()

	handlers := &Handlers{
		connection: conn,
	}

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/api/store", handlers.StoreHandler)
	http.ListenAndServe(":8080", nil)
}

type Handlers struct {
	connection rediscli.Conn
}

func (h *Handlers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world.")
}

func (h *Handlers) StoreHandler(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	if _, err := h.connection.Do("SET", key, value); err != nil {
		fmt.Fprintf(w, "Error : %s", err.Error())
		return
	}

	fmt.Fprintf(w, "Key : %s, Value : %s", key, value)
}
