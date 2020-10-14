package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/env", Env)
	http.HandleFunc("/health", HealthCheck)
	http.ListenAndServe(":8080", nil)
}

//HealthCheck - return 200
func HealthCheck(w http.ResponseWriter, r *http.Request) {}

//Env - print current env vars
func Env(w http.ResponseWriter, r *http.Request) {
	for _, s := range os.Environ() {
		fmt.Fprintf(w, "%s\n", s)
	}
}
