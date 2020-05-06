package main

import (
	"fmt"
	"headers/codes"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var defaultHTPStatus int = 200
var hostname string
var port string

func init() {
	var err error
	var ret bool

	hostname, err = os.Hostname()
	if err != nil {
		panic(err)
	}

	port, ret := os.LookupEnv("PORT")
	if !ret {
		port = ":9090"
	} else {
		port = ":" + port
	}

}

func main() {

	log.Printf("Starting Headers on port: %s\n", port)

	http.HandleFunc("/", printHeaders)
	http.HandleFunc("/set/", setStatusCode)
	http.ListenAndServe(port, nil)
}

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
		} else if !codes.CheckHTTPCODE(intCode) {
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
