package hashid

import (
	"fmt"
	"strings"
	"unicode"
)

// UrlParam returns urlparam
func UrlParam(title string, pk int64) string {
	var sb strings.Builder

	wordCount, maxWordCount := 0, 10
	runeCount, maxRuneCount := 0, 300
	lastChar, hyphen := rune('-'), rune('-')
	for _, r := range []rune(title) {
		if runeCount == maxRuneCount || wordCount == maxWordCount {
			break
		}
		// replace space in title with hyphen while making sure
		// consequent hyphens do not occur
		if unicode.IsSpace(r) {
			if wordCount == maxWordCount-1 {
				break
			}
			if lastChar != hyphen {
				sb.WriteRune(hyphen)
				wordCount += 1
				runeCount += 1
			}
			lastChar = rune('-')
			continue
		}
		// retain hyphen from title while making sure
		// consequent hyphens do not occur
		if r == hyphen {
			if lastChar != hyphen {
				sb.WriteRune(hyphen)
				runeCount += 1
				lastChar = rune(hyphen)
			}
			continue
		}
		// retain all language alphabet and numbers from title
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			sb.WriteRune(unicode.ToLower(r))
			runeCount += 1
			lastChar = r
			continue
		}
	}

	return fmt.Sprintf(
		"%s-%s",
		sb.String(), Encode(pk),
	)
}
