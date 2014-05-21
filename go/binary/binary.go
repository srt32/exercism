package binary

import (
  sc "strconv"
  s "strings"
)

func ToDecimal(binary string) (decimal int) {
  digits := s.Split(binary, "")
  for i, v := range digits {
    switch {
    case v > sc.Itoa(1):
      return 0
    }

    num, _ := sc.Atoi(v)
    amount := num << uint(len(digits)-i-1)
    decimal += int(amount)
  }
  return decimal
}
