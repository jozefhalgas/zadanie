package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	requestCounter int
	hostname       string
)

/* Check if app is healthy or not */
func CheckAppHealth() (int, string) {

	if _, err := os.Stat("/tmp/fail-health-check"); err == nil {
		return http.StatusInternalServerError, "BAD"
	}

	return http.StatusOK, "GOOD"
}

func handler(w http.ResponseWriter, r *http.Request) {
	if code, _ := CheckAppHealth(); code != 500 {
		requestCounter++
		fmt.Fprintf(w, "Hi there, got %d requests, running on %s, I love %s!", requestCounter, hostname, r.URL.Path[1:])
	} else {
		w.WriteHeader(code)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	var HcStatus string

	Hccode, HcStatus := CheckAppHealth()

	w.WriteHeader(Hccode)
	fmt.Fprintf(w, "This is - %s - health check... at %s/%s", HcStatus, hostname, r.URL.Path[1:])
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	requestCounter = 0
	hostname = os.Getenv("HOSTNAME")
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", health)
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8080", nil)
}
