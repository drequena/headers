package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestPrintHeaders(t *testing.T) {

	httpWantedCode := 200
	httpPatternWantedBody := "{\"STATUS:\" \"200\", \"HOST:\" \".*\"}\n"

	status, body := makeTestRequest("GET", "http://127.0.0.1/set/500", nil, printHeaders)

	if status != httpWantedCode {
		t.Errorf("Expected StatusCode %d, got %d", httpWantedCode, status)
	}

	regexRet, err := regexp.MatchString("{\"STATUS:\" \"200\", \"HOST:\" \".*\"}", body)
	if !regexRet || err != nil {
		t.Errorf("Expected StatusCode pattern %s, got %s", httpPatternWantedBody, body)
	}
}

func TestSetStatusCode(t *testing.T) {

	httpWantedCode := 200
	httpWantedBody := "{\"STATUS:\" \"OK\"}\n"

	status, body := makeTestRequest("GET", "http://127.0.0.1/set/500", nil, setStatusCode)

	if status != httpWantedCode {
		t.Errorf("Expected StatusCode %d, got %d", httpWantedCode, status)
	}

	if httpWantedBody != body {
		t.Errorf("Expected StatusCode pattern %s, got %s", httpWantedBody, body)
	}

}

func TestSetStatusCodeInvalidHTTPCode(t *testing.T) {
	httpWantedCode := 500
	httpWantedBody := "{\"STATUS:\" \"Invalid HTTP Code\"}\n"

	status, body := makeTestRequest("GET", "http://127.0.0.1/set/700", nil, setStatusCode)

	if status != httpWantedCode {
		t.Errorf("Expected StatusCode %d, got %d", httpWantedCode, status)
	}

	if httpWantedBody != body {
		t.Errorf("Expected StatusCode pattern %s, got %s", httpWantedBody, body)
	}
}

func TestSetStatusCodeConvertError(t *testing.T) {
	httpWantedCode := 500
	httpWantedBody := "{\"STATUS:\" \"Error\"}\n"

	status, body := makeTestRequest("GET", "http://127.0.0.1/set/abcde", nil, setStatusCode)

	if status != httpWantedCode {
		t.Errorf("Expected StatusCode %d, got %d", httpWantedCode, status)
	}

	if httpWantedBody != body {
		t.Errorf("Expected StatusCode pattern %s, got %s", httpWantedBody, body)
	}
}

func TestSetStatusCodeWrongMethod(t *testing.T) {
	httpWantedCode := http.StatusMethodNotAllowed
	httpWantedBody := ""

	status, body := makeTestRequest("POST", "http://127.0.0.1/set/abcde", nil, setStatusCode)

	if status != httpWantedCode {
		t.Errorf("Expected StatusCode %d, got %d", httpWantedCode, status)
	}

	if httpWantedBody != body {
		t.Errorf("Expected StatusCode pattern %s, got %s", httpWantedBody, body)
	}
}

//TestCheckHTTPCODE test if HTTPCODE is correct
func TestCheckHTTPCODE(t *testing.T) {

	var got bool
	want := true

	for code := range validCodes {
		got = CheckHTTPCODE(code)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	got = CheckHTTPCODE(1)
	if got == want {
		t.Errorf("got %v want %v", got, want)
	}
}

func makeTestRequest(method string, url string, body io.Reader, handler func(http.ResponseWriter, *http.Request)) (statusCode int, stringBody string) {
	request := httptest.NewRequest(method, url, body)
	response := httptest.NewRecorder()
	request.Header.Add("Tests", "goTests")
	handler(response, request)
	retBody, _ := ioutil.ReadAll(response.Body)

	return response.Result().StatusCode, string(retBody)
}

// func TestPrintHeaders(t *testing.T) {

// 	httpWantedCode := 200
// 	httpPatternWantedBody := "{\"STATUS:\" \"200\", \"HOST:\" \".*\"}\n"

// 	r := httptest.NewRequest("GET", "http://127.0.0.1/", nil)
// 	w := httptest.NewRecorder()
// 	r.Header.Add("Test", "OK")
// 	printHeaders(w, r)
// 	resp := w.Result()
// 	body, _ := ioutil.ReadAll(resp.Body)

// 	if resp.StatusCode != httpWantedCode {
// 		t.Errorf("Expected StatusCode %d, got %d", httpWantedCode, resp.StatusCode)
// 	}

// 	regexRet, err := regexp.MatchString("{\"STATUS:\" \"200\", \"HOST:\" \".*\"}", string(body))
// 	if !regexRet || err != nil {
// 		t.Errorf("Expected StatusCode pattern %s, got %s", httpPatternWantedBody, string(body))
// 	}
// }

// func TestSetStatusCode(t *testing.T) {

// 	httpWantedCode := 200
// 	httpWantedBody := "{\"STATUS:\" \"OK\"}\n"

// 	r := httptest.NewRequest("GET", "http://127.0.0.1/set/500", nil)
// 	w := httptest.NewRecorder()
// 	setStatusCode(w, r)
// 	resp := w.Result()
// 	body, _ := ioutil.ReadAll(resp.Body)

// 	if resp.StatusCode != httpWantedCode {
// 		t.Errorf("Expected StatusCode %d, got %d", httpWantedCode, resp.StatusCode)
// 	}

// 	if httpWantedBody != string(body) {
// 		t.Errorf("Expected StatusCode pattern %s, got %s", httpWantedBody, string(body))
// 	}

// }
