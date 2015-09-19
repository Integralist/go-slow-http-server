package main

import (
	"fmt"
	"net/http"
	"time"
)

func Seconds(n int) time.Duration {
	return time.Duration(n) * time.Second
}

func handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	time.Sleep(Seconds(5))
	elapsed := time.Since(now)
	fmt.Fprintf(w, "Request to /%s took %s", r.URL.Path[1:], elapsed)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
