package queenattack

import (
	"errors"
	"math"
	"strconv"
)

func CanQueenAttack(w string, b string) (attack bool, err error) {
	if (w == b) {
		return false, errors.New("Same square")
	}

	white := position{w}
	black := position{b}

	boundsError := checkBounds(white, black)
	if boundsError != nil {
		return false, boundsError
	}

	if white.File() == black.File() {
    return true, nil
	} else if white.Rank() == black.Rank() {
    return true, nil
  } else if areDiagonal(white, black) {
		return true, nil
	} else {
	  return false, nil
	}
};

func areDiagonal(w, b position) bool {
	fileDelta := float64(b.FileNumber() - w.FileNumber())
	rankDelta := float64(b.Rank() - w.Rank())
	if (math.Abs(fileDelta) == math.Abs(rankDelta)) {
		return true
	} else {
		return false
	}
}

type position struct {
	Location string
}

func (p position) File() string {
	 return string(p.Location[0])
}

func (p position) FileNumber() int64 {
	return map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}[p.File()]
}

func (p position) Rank() int64 {
	rank, err := strconv.ParseInt(string(p.Location[1]), 0, 32)
	if err != nil {
		return 9 // hacky
	}
	return rank
}

func checkBounds(white, black position) error {
	if white.outOfBounds() {
		return errors.New("Out of bounds")
	} else if black.outOfBounds() {
		return errors.New("Out of bounds")
	} else {
		return nil
	}
}

func (p position) outOfBounds() bool {
	if p.Rank() > 8 {
	  return true
	} else {
		return false
	}
}
