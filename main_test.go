package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

// Source https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
var validCodes = map[int]bool{
	100: true,
	101: true,
	102: true,
	103: true,
	200: true,
	202: true,
	203: true,
	204: true,
	205: true,
	206: true,
	207: true,
	208: true,
	226: true,
	300: true,
	301: true,
	302: true,
	303: true,
	304: true,
	305: true,
	307: true,
	308: true,
	400: true,
	401: true,
	402: true,
	403: true,
	404: true,
	405: true,
	406: true,
	407: true,
	408: true,
	409: true,
	410: true,
	411: true,
	412: true,
	413: true,
	414: true,
	415: true,
	416: true,
	417: true,
	421: true,
	422: true,
	423: true,
	424: true,
	425: true,
	426: true,
	428: true,
	429: true,
	431: true,
	451: true,
	500: true,
	501: true,
	502: true,
	503: true,
	504: true,
	505: true,
	506: true,
	507: true,
	508: true,
	510: true,
	511: true}

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
