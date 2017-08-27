package triangle

import "math"

const testVersion = 3

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func isTriangle(a, b, c float64) bool {
	return true &&
		a > 0 && a < math.Inf(1) && // a valid length
		b > 0 && b < math.Inf(1) && // b valid length
		c > 0 && c < math.Inf(1) && // c valid length
		(a+b) >= c && // sides much add up
		(a+c) >= b &&
		(b+c) >= a
}

func sidesMatch(a, b float64) bool {
	return a == b
}

func KindFromSides(a, b, c float64) Kind {
	kind := Sca
	if !isTriangle(a, b, c) {
		kind = NaT
	} else if sidesMatch(a, b) && sidesMatch(a, c) {
		kind = Equ
	} else if sidesMatch(a, b) || sidesMatch(a, c) || sidesMatch(b, c) {
		kind = Iso
	}
	return kind
}
