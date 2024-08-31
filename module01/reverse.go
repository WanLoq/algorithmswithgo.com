package module01

import (
	"slices"
	"strings"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	asArray := strings.Split(word, "")
	slices.Reverse(asArray)
	output := strings.Join(asArray, "")
	return output
}
