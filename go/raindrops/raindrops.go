package raindrops

import (
  "strconv"
)

func Convert(n int) string {
  primes := map[int]string{
    3: "Pling",
    5: "Plang",
    7: "Plong",
  }

  result := createResult(n, primes)

  if result == "" {
    return strconv.Itoa(n)
  } else {
    return result
  }
}

func createResult(n int, primes map[int]string) (result string) {
  for prime, saying := range primes {
    if n%prime == 0 {
      result += saying
    }
  }
  return result
}
