package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {

	http.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hey from go on %s!", runtime.GOOS)
	})

	http.HandleFunc("/find", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Searching for %s!", r.FormValue("search"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../../../www/find.html")
	})

	log.Fatal(http.ListenAndServe(":4269", nil))
}
