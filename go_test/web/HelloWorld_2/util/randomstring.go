package util

import (
	"math/rand"
	"time"
)

const letterBytes = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func RandomString(n int) string {
	res := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range res {
		res[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(res)
}
