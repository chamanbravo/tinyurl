package utils

import (
	"math/rand"
	"strings"
)

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func UrlGenerator(n int) string {
	b := make([]string, n)
	for i := range b {
		b[i] = string(letters[rand.Intn(len(letters))])
	}
	randString := strings.Join(b, "")
	return randString
}
