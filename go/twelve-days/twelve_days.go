package twelve

import (
	_ "fmt"
	_ "strings"
)

const testVersion = 1

type Day struct {
	day  string
	gift string
}

var days = []Day{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

func addGifts(line string, day int) string {
	for i := day - 1; i > 0; i-- {
		//line = fmt.Sprintf("%s, %s", line, days[i].gift)
		//line = strings.Join([]string{line, days[i].gift}, ", ")
		line = line + ", " + days[i].gift
	}
	if day > 1 {
		//line = fmt.Sprintf("%s, and %s.", line, days[0].gift)
		//line = strings.Join([]string{line, ", and ", days[0].gift, "."}, "")
		line = line + ", and " + days[0].gift + "."
	} else {
		//line = fmt.Sprintf("%s, %s.", line, days[0].gift)
		//line = strings.Join([]string{line, ", ", days[0].gift, "."}, "")
		line = line + ", " + days[0].gift + "."
	}
	return line
}

func Verse(day int) string {
	//line := fmt.Sprintf("On the %s day of Christmas my true love gave to me", days[day-1].day)
	//line := strings.Join([]string{"On the", days[day-1].day, "day of Christmas my true love gave to me"}, " ")
	line := "On the " + days[day-1].day + " day of Christmas my true love gave to me"
	return addGifts(line, day)
}

func Song() (song string) {
	for i, _ := range days {
		//song = fmt.Sprintf("%s%s\n", song, Verse(i+1))
		//song = strings.Join([]string{song, Verse(i + 1), "\n"}, "")
		song = song + Verse(i+1) + "\n"
	}
	return song
}
