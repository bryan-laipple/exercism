package clock

import "fmt"

const testVersion = 4

type Clock interface {
	String() string
	Add(minutes int) Clock
}

type clockImpl struct {
	minutes int
}

// returns an implementation of Clock
func New(hour, minute int) Clock {
	clock := clockImpl{0}
	clock.minutes += hour * 60
	clock.minutes += minute
	clock.normalize()
	return clock
}

func (clock *clockImpl) normalize() {
	minutes := clock.minutes % 60
	hours := clock.minutes / 60
	hours = hours % 24
	clock.minutes = (hours * 60) + minutes
	for clock.minutes < 0 {
		clock.minutes += 24 * 60
	}
}

func (clock clockImpl) String() string {
	minute := clock.minutes % 60
	hour := clock.minutes / 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

func (clock clockImpl) Add(minutes int) Clock {
	clock.minutes += minutes
	clock.normalize()
	return clock
}
