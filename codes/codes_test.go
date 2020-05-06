package codes

import (
	"testing"
)

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
