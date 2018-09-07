package utils

import (
	"math/rand"
	"time"
)

func RandNumber(min,max int64) int64{
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}
