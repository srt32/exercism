package clock

import (
	"strconv"
)

func New(h, m int) clock {
	return clock{h, m}
}

type clock struct {
	h int
	m int
}

func (c clock) String() string {
	minutes := strconv.FormatInt(int64(c.m), 10)
	hours := strconv.FormatInt(int64(c.h), 10)

	if c.m == 0 {
		minutes += "0"
	} else if c.m < 10 {
		minutes = "0" + minutes
	}

	if c.h < 10 {
		hours = "0" + hours
	}

	display := hours + ":" + minutes

	return display
}

func (c clock) Add(a int) clock {
	hours := c.h
	minutes := c.m
	newMinutes := c.m + a

	if newMinutes > 59 {
		hours += 1
		minutes -= 60
	} else {
		minutes = newMinutes
	}

	return clock{hours, minutes}
}
