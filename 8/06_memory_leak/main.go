package main

import (
	"net/http"
	_ "net/http/pprof"
)

func WaitingFunc(c chan struct{}) {
	<-c
}

func Handler(w http.ResponseWriter, r *http.Request) {
	c := make(chan struct{}, 3)
	go WaitingFunc(c)

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", Handler)

	http.ListenAndServe(":8080", nil)
}


