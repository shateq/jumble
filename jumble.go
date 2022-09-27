// Provides jumbling-ciphering methods
package jumble

import (
	"math/rand"
	"strings"
)

// Shuffle every middle letters of every word in `text` string
func Jumble(text string) string {
	fields := strings.Fields(text)
	var out []string

	for _, el := range fields {
		if len(el) <= 3 {
			out = append(out, el)
			continue
		}

		f, l, rest := string(el[0]), string(el[len(el)-1]), string(el[1:len(el)-1])

		runes := []rune(rest)
		rand.Shuffle(len(runes), func(i, j int) {
			runes[i], runes[j] = runes[j], runes[i]
		})

		rest = f + string(runes) + l
		out = append(out, rest)
	}

	value := strings.Join(out, " ") + " "
	if strings.Contains(text, "\n") {
		value = value + "\n"
	}

	return value
}

// Swaps two middle letters of every word in `text` string
func MixTwo(text string) string {
	// TODO: implement placeholders
	// TODO: create algorithm
	return "TODO"
}
