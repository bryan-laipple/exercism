package bob

import (
	re "regexp"
	s "strings"
)

const testVersion = 3

func isShouting(input *string) bool {
	upper := s.ToUpper(*input)
	lower := s.ToLower(*input)
	return upper != lower && upper == *input
}

func isQuestion(input *string) bool {
	return s.LastIndex(*input, "?") == len(*input)-1
}

func isSilence(input *string) bool {
	match, _ := re.MatchString("^\\s+$", *input)
	return match || len(*input) == 0
}

func Hey(input string) (output string) {
	input = s.Trim(input, " ")
	if isSilence(&input) {
		output = "Fine. Be that way!"
	} else if isShouting(&input) {
		output = "Whoa, chill out!"
	} else if isQuestion(&input) {
		output = "Sure."
	} else {
		output = "Whatever."
	}
	return
}
