package raindrops

import (
  "strconv"
  "unicode/utf8"
)

func Convert(n int) string {
  values := map[int]string{
    3: "Pling",
    5: "Plang",
    7: "Plong",
  }

  result := ""

  for k, v := range values {
    if n%k == 0 {
      result = result + v
    }
  }

  if utf8.RuneCountInString(result) == 0 {
    return strconv.Itoa(n)
  } else {
    return result
  }
}
