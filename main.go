package main

import (
	"log"
	"net/http"
	"os"
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

	port, ret = os.LookupEnv("PORT")
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
