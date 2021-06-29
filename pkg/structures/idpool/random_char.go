package idpool

import (
	"math/rand"
	"strings"
)

var alpha = "abcdefghijklmnopqrstuvwxyz"
var alphanumeric = "0123456789" + alpha + strings.ToUpper(alpha)

func generateChar() uint8 {
	return alphanumeric[rand.Intn(len(alphanumeric))]
}
