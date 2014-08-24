package queenattack

import (
	"errors"
	"strconv"
)

func CanQueenAttack(w string, b string) (attack bool, err error) {
	white := position{w}
	black := position{b}

	if (w == b) {
		return false, errors.New("Same square")
	} else if white.outOfBounds() {
		return false, errors.New("Out of bounds")
	} else if black.outOfBounds() {
		return false, errors.New("Out of bounds")
	} else {
		// b4, b7, handle same file
	  return false, nil
	}
};

type position struct {
	Location string
}

func (p position) outOfBounds() bool {
	location := p.Location
	rank, err := strconv.ParseInt(string(location[1]), 0, 32)
	if err != nil {
		return true
	}
	if rank > 8 {
	  return true
	} else {
		return false
	}
}
