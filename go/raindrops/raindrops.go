package raindrops

import (
	"fmt"
	s "strings"
)

const testVersion = 3

func checkFactor(in int, factor int, raindrop string, out string) string {
	if in%factor == 0 {
		out = s.Join([]string{out, raindrop}, "")
	}
	return out
}

func Convert(in int) (out string) {
	out = checkFactor(in, 3, "Pling", out)
	out = checkFactor(in, 5, "Plang", out)
	out = checkFactor(in, 7, "Plong", out)

	if len(out) == 0 {
		out = fmt.Sprintf("%d", in)
	}
	return
}
