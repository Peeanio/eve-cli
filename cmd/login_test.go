package cmd

import (
	"testing"
)

func Test_randomBytesInHex(t *testing.T) {
	result, err := randomBytesInHex(3)
	if err != nil {
		t.Error("incorrect result: got error from function", err)
	} else if len(result) != 6 {
		t.Error("incorrect result: expected length 6, got", len(result))
	}
}

func Test_authorizationURL(t *testing.T) {
	codeVerifier, codeChallenge, state := authorizationURL()
	if len(codeVerifier) != 64 {
		t.Error("incorrect result: codeVerifier length expected 64", len(codeVerifier))
	} else if len(codeChallenge) != 43 {
		t.Error("incorrect result: codeChallenge lengthexpected 43", len(codeChallenge))
	} else if len(state) != 48 {
		t.Error("incorrect result: state lengthexpected 48", len(state))
	}
}
