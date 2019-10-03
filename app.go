package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type SampleJSONResponse struct {
	Message string
        Path    string
}

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!, call to endpoint %s!", r.URL.Path[1:])
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getJSONresponse(w http.ResponseWriter, r *http.Request) {
        res := SampleJSONResponse{"A JSON message", r.URL.Path[1:]}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
        http.HandleFunc("/", getJSONresponse)
	http.HandleFunc("/hello", getHelloWorld)
	http.HandleFunc("/health", healthCheck)
	http.ListenAndServe(":8080", nil)
}
