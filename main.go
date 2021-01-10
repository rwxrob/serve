package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func serveForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	jsn, _ := json.MarshalIndent(r.Form, "", "  ")
	fmt.Fprintf(w, string(jsn))
}

func serve(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		serveForm(w, r)
		return
	}
	var err error
	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}
	http.ServeFile(w, r, wd+r.URL.Path)
}

func main() {
	http.HandleFunc("/", serve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
