package mask

import (
	"math/rand"
	"time"
	"unicode"
)

func CharMapping(c rune, replace rune,fixed bool) rune {
	rand.Seed(time.Now().UnixNano())
	if unicode.IsDigit(c) {
		return rune('0'+rand.Intn(10))
	} else if unicode.IsLetter(c) {
		if(fixed) {
			return replace
		}
		if c <= 90 {
			return rune('A'+rand.Intn(26))
		} else {
			return rune('a'+rand.Intn(26))
		}
	}
	return replace
}
