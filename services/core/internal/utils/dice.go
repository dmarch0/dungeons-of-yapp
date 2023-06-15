package utils

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}

func ND6(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		result += D6()
	}

	return result
}

func D6() int {
	return r.Intn(5) + 1
}
