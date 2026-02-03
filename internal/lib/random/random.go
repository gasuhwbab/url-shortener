package random

import (
	"math/rand"
	"time"
)

func NewRandomStrinng(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	alph := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, size)
	for i := range b {
		b[i] = alph[rnd.Intn(len(alph))]
	}
	return string(b)
}
