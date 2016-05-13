package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	sections := strings.Split(r.URL.Path, "/")
	if len(sections) > 2 {
		delay, err := strconv.Atoi(sections[2])
		if err != nil {
			fmt.Fprintf(w, "Hmm, looks like there was a problem with your URL")
		}

		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Fprintf(w, "I love %s! (delayed by %d seconds)", sections[1], delay)
		return
	}

	if len(sections) > 1 {
		item := r.URL.Path[1:]
		if item != "" {
			fmt.Fprintf(w, "I love %s! (no delay)", item)
		} else {
			fmt.Fprintf(w, "Usage example /foo/5")
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
