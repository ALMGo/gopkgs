package randstring

import "testing"

const strLen = 4

func TestRandStringWithCharset(t *testing.T) {
	str := RandStringWithCharset(strLen, "a")

	if len(str) != strLen {
		t.Fatalf("RandStringWithCharset returns incorrect length string")
	}

	if str != "aaaa" {
		t.Fatalf("RandStringWithCharset did not return expected string")
	}
}

func TestRandString(t *testing.T) {
	str := RandString(strLen)

	if len(str) != strLen {
		t.Fatalf("RandString returns incorrect length string")
	}
}
