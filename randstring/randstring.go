package randstring

import (
	"math/rand"
	"time"
)

const CapitalLettersCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LowerLettersCharset = "abcdefghijklmnopqrstuvwxyz"
const NumbersCharset = "0123456789"
const LettersCharset = CapitalLettersCharset + LowerLettersCharset
const DefaultCharset = LettersCharset + NumbersCharset

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return RandomStringWithCharset(length, DefaultCharset)
}
