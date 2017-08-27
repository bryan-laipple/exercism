package accumulate

const testVersion = 1

func Accumulate(before []string, op func(string) string) (after []string) {
	for _, s := range before {
		after = append(after, op(s))
	}
	return
}
