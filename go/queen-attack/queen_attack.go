package queenattack

import (
	"errors"
)

func CanQueenAttack(w string, b string) (attack bool, err error) {
	if (w == b) {
		err := errors.New("Same square")
		return false, err
	} else {
	  return false, nil
	}
};
