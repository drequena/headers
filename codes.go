package main

import "net/http"

// CheckHTTPCODE check if a int is a valid HTTP STATUS CODE
func CheckHTTPCODE(code int) bool {

	if http.StatusText(code) != "" {
		return true
	}

	return false
}
