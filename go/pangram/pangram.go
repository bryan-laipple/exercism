package pangram

import "strings"

const testVersion = 2

func letters() map[rune]bool {
	letters := make(map[rune]bool)
	for i := 0; i < 26; i++ {
		letters[rune('a'+i)] = false
	}
	return letters
}

func IsPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)
	letters := letters()
	for _, c := range sentence {
		delete(letters, c)
	}
	return len(letters) == 0
}
