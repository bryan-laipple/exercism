package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	divisibleBy4 := year%4 == 0
	divisbleBy100 := year%100 == 0
	divisbleBy400 := year%400 == 0
	return divisibleBy4 && (!divisbleBy100 || divisbleBy400)
}
