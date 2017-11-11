package twelve

import (
	"bytes"
)

const testVersion = 1

func Verse(day int) string {
	var items, daysNth = Verses()
	var dayNth = daysNth[day-1]
	var buffer bytes.Buffer
	var onThe = "On the " + dayNth + " day of Christmas my true love gave to me, "
	buffer.WriteString(onThe)
	if day == 1 {
		buffer.WriteString(items[0] + ".")
	} else {
		for day > 1 {
			var arrayIndex = day - 1
			var item = items[arrayIndex]
			buffer.WriteString(item + ", ")
			day--
		}
		buffer.WriteString("and " + items[0] + ".")
	}
	return buffer.String()
}

func Song() string {
	var buffer bytes.Buffer
	var day = 1
	for day <= 12 {
		buffer.WriteString(Verse(day) + "\n")
		day++
	}
	return buffer.String()
}

func Verses() ([12]string, []string) {
	var items [12]string
	items[0] = "a Partridge in a Pear Tree"
	items[1] = "two Turtle Doves"
	items[2] = "three French Hens"
	items[3] = "four Calling Birds"
	items[4] = "five Gold Rings"
	items[5] = "six Geese-a-Laying"
	items[6] = "seven Swans-a-Swimming"
	items[7] = "eight Maids-a-Milking"
	items[8] = "nine Ladies Dancing"
	items[9] = "ten Lords-a-Leaping"
	items[10] = "eleven Pipers Piping"
	items[11] = "twelve Drummers Drumming"
	daysNth := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	return items, daysNth
}
