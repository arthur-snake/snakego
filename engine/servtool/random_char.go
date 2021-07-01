package servtool

import (
	"math/rand"
	"strings"
)

var alpha = "abcdefghijklmnopqrstuvwxyz"
var alphanumeric = "0123456789" + alpha + strings.ToUpper(alpha)

func generateChar() uint8 {
	// TODO: use better random?
	return alphanumeric[rand.Intn(len(alphanumeric))] //nolint:gosec
}
