package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var defaultHTPStatus int = 200

func main() {
	port, err := os.LookupEnv("PORT")
	if !err {
		port = ":9090"
	} else {
		port = ":" + port
	}

	http.HandleFunc("/", printHeaders)
	http.HandleFunc("/set/", setStatusCode)
	http.ListenAndServe(port, nil)
}

func printHeaders(w http.ResponseWriter, r *http.Request) {

	for x, z := range r.Header {
		fmt.Printf("%s - %s\n", x, z)
	}

	fmt.Printf("Path: %s\nMethod: %s\n", r.URL.Path, r.Method)

	//buf := new(bytes.Buffer)
	// buf.ReadFrom(r.Body)
	// fmt.Printf("Body: %s", buf.String())

	hostname, _ := os.Hostname()
	w.WriteHeader(defaultHTPStatus)
	fmt.Fprintf(w, "{\"STATUS:\" \"OK\", \"HOST:\" \"%s\"}\n", hostname)
}

func setStatusCode(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		strCode := strings.TrimPrefix(r.URL.Path, "/set/")
		intCode, err := strconv.Atoi(strCode)

		if err != nil {
			fmt.Printf("Error converting %s to int", strCode)
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error\n")
		} else {
			fmt.Printf("Changing / status code to: %d\n", intCode)
			defaultHTPStatus = intCode
			fmt.Fprintf(w, "{\"STATUS:\" \"OK\"}\n")
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "")
	}
}
