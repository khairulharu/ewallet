package util

import "math/rand"

func GenerateRandomString(n int) string {
	var charsets = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	letters := make([]rune, n)

	for i := range letters {
		letters[i] = charsets[rand.Intn(len(charsets))]
	}

	return string(letters)
}
