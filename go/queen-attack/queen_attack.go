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
	//} else if sameFile(white, black) {
	// create file function and then use it for same file check

	} else {
		// b4, b7, handle same file
	  return false, nil
	}
};

//func sameFile(w, b position) {
//	if
//}

type position struct {
	Location string
}

func (p position) rank() int64 {
	rank, err := strconv.ParseInt(string(p.Location[1]), 0, 32)
	if err != nil {
		return 9 // hacky
	}
	return rank
}

func (p position) outOfBounds() bool {
	if p.rank() > 8 {
	  return true
	} else {
		return false
	}
}
