package rand

import (
	"math/rand"
	"strings"
)

func String() string {
	i := rand.Intn(62)
	if i < 10 {
		return string(48 + i)
	}
	if i < 36 {
		return string(55 + i)
	}
	return string(61 + i)
}

func Stringn(n int) string {
	s := make([]string, n)
	for i := 0 ; i < n ; i++ {
		s[i] = String()
	}
	return strings.Join(s, "")
}
