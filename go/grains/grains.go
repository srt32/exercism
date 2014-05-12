package grains

import (
  "math"
)

func Square(number int) uint64 {
  result := math.Exp2(float64(number - 1))
  return uint64(result)
}

func Total() uint64 {
  return Square(65) - 1
}
