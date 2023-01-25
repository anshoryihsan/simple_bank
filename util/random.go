package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstufwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInit generate a random integer betwheen min and max
func RandomInit(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //max-min
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len((alphabet))

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInit(0, 100)
}

func RandomCurrency() string {
	currencies := []string{"IDR", "USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
