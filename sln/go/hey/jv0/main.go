package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hey from go on %s!", runtime.GOOS)
	})

	log.Fatal(http.ListenAndServe(":4269", nil))
}
