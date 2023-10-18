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
