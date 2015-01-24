package clock

import (
	"fmt"
	"math"
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
	var hour = strconv.Itoa(c.h)
	var minute = strconv.Itoa(int(math.Abs(float64(c.m))))

	if c.h < 10 {
		hour = "0" + hour
	}

	if c.m < 10 && c.m > -1 {
		minute = "0" + minute
	}

	return hour + ":" + minute
}

func (c clock) Add(duration int) clock {
	var newHour = c.h
	var newMinute = c.m

	var numHours = duration / 60
	newHour = newHour + numHours

	newMinute = c.m + (duration - 60*numHours)

	fmt.Println(newMinute)

	if newMinute > 59 {
		newMinute = newMinute - 60
	}

	if newMinute < 0 {
		newHour -= 1
	}

	if newHour > 23 {
		newHour = c.h - 24
	}
	return clock{newHour, newMinute}
}
