package main

import (
	"log"
	"net/http"
	"time"
)

type LifeCycle struct {
	shutdownWaitTimeInSeconds int
}

func (l *LifeCycle) PostStartHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PostStartHandler called")
	w.WriteHeader(200)
}

func (l *LifeCycle) PreStopHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PreStopHandler : called")
	log.Print("PreStopHandller : starting wait")

	time.Sleep(time.Duration(l.shutdownWaitTimeInSeconds) * time.Second)

	log.Print("PreStopHandller : bye...")

	w.WriteHeader(200)
}

func main() {

	lifecycle := &LifeCycle{
		shutdownWaitTimeInSeconds: 10,
	}

	http.HandleFunc("/lifecycle/poststart", lifecycle.PostStartHandler)
	http.HandleFunc("/lifecycle/prestop", lifecycle.PreStopHandler)
	http.ListenAndServe(":8080", nil)
}
