package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../../../www/find.html")
	})

	http.HandleFunc("/find", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Searching for %s!", r.FormValue("search"))
	})

	log.Fatal(http.ListenAndServe(":4269", nil))
}
