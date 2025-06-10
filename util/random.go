package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		s := alphabet[rand.Intn(k)]
		sb.WriteByte(s)
	}

	return sb.String()
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomCurrency() string {
	c := []string{"USD", "EUR", "RUB"}
	return c[rand.Intn(len(c))]
}
