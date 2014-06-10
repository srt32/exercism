package diffsquares

import(
  . "math"
)

func SquareOfSums(n int) (sum int) {
  for i := 0; i <= n; i++ {
    sum += i
  }
  return int(Pow(float64(sum), 2))
}

func SumOfSquares(n int) (sum int) {
  for i := 0; i <= n; i++ {
    sum += int(Pow(float64(i), 2))
  }
  return
}

func Difference(n int) int {
  return SquareOfSums(n) - SumOfSquares(n)
}
