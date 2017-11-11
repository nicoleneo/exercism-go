package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	for minute < 0 {
		minute += 60
		hour -= 1
	}
	minutes := minute % 60
	hours := (hour + (minute / 60)) % 24
	for hours < 0 {
		hours += 24
	}

	return Clock{hours, minutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	hours := minutes / 60
	minutes = minutes % 60
	return New(c.hour+hours, c.minute+minutes)
}
