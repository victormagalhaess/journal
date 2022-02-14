package art

import (
	"strings"
	"time"
)

type Time struct {
	Time time.Time
}

func BuildAsciiArt(input string) string {
	art := []string{"0", "/", "&", "o", "#", "%", "O", "8", "@", "*"}

	if len(input) == 0 {
		return ""
	}

	for _, c := range input {
		if c < '0' || c > '9' {
			return ""
		}
	}

	var result strings.Builder
	for _, c := range input {

		index := int(c) % len(art)
		result.WriteString(art[index])
	}
	return result.String()
}
