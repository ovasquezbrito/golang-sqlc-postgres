package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt Genera numeros aleatorios entre un mínimo u un  maximo
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString genera string aleatorios
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner .Gebera nombres aleatorios para el campo owner
func RandomOwner() string {
	return RandomString(6)
}

//RandomMeney Genera monedas aleatorias del 0 al 1000
func RandomMeney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency general currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
