package triangle

import(
  "math"
)

type Kind interface {}

const Equ = "equ"
const Iso = "iso"
const Sca = "sca"
const NaT = "nat"

func KindFromSides(a, b, c float64) Kind {
  for _, side := range []float64{a, b, c} {
    if math.IsNaN(side) || math.IsInf(side, 1) || math.IsInf(side, -1) {
      return NaT
    }
  }
  if a + b <= c || a + c <= b || b + c <= a {
    return NaT
  }
  if a == b && b == c {
    return Equ
  } else if a == b || b == c || a == c {
    return Iso
  }
  return Sca
}
