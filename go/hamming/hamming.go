package hamming

import (
	"errors"
	"strings"
)

const testVersion = 6

func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		distance = -1
		err = errors.New("DNA strands not equal length.")
		return
	}
	a = strings.ToUpper(a)
	b = strings.ToUpper(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance += 1
		}
	}
	return
}
