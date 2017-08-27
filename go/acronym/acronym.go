package acronym

import s "strings"

const testVersion = 3

func Abbreviate(in string) (acronym string) {
	letters := make([]rune, 0, len(in))
	startOfWord := true
	for _, c := range in {
		if c == ' ' || c == '-' {
			startOfWord = true
			continue
		}
		if startOfWord {
			letters = append(letters, c)
			startOfWord = false
		}
	}
	acronym = string(letters)
	acronym = s.ToUpper(acronym)
	return

}
