package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func printHeaders(w http.ResponseWriter, r *http.Request) {

	for x, z := range r.Header {
		fmt.Printf("%s - %s\n", x, z)
	}

	fmt.Printf("Path: %s\nMethod: %s\n", r.URL.Path, r.Method)

	w.WriteHeader(defaultHTPStatus)
	fmt.Fprintf(w, "{\"STATUS:\" \"%d\", \"HOST:\" \"%s\"}\n", defaultHTPStatus, hostname)
}

func setStatusCode(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		strCode := strings.TrimPrefix(r.URL.Path, "/set/")
		intCode, err := strconv.Atoi(strCode)

		if err != nil {
			log.Printf("Error converting %s to int", strCode)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"STATUS:\" \"Error\"}\n")
		} else if !CheckHTTPCODE(intCode) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{\"STATUS:\" \"Invalid HTTP Code\"}\n")
		} else {
			log.Printf("Changing / status code to: %d\n", intCode)
			defaultHTPStatus = intCode
			fmt.Fprintf(w, "{\"STATUS:\" \"OK\"}\n")
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "")
	}
}
