package main

import (
	"net/http"
	_ "net/http/pprof"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

// cd $GOPATH/src/github.com/uber/go-torch
// git clone https://github.com/brendangregg/FlameGraph.git

