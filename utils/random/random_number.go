package random

import (
	"math/rand"
	"time"
)

// GetRandom ...
func GetRandom() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(99999999999)
}
