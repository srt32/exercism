package secret

import (
  "math"
  "strconv"
)

func Handshake(code int) []string {
  if (code < 0 || code > 31) {
    return []string{}
  }

  binaryCode := strconv.FormatInt(int64(code), 2)

  codes := map[int]string{
    1: "wink",
    2: "double blink",
    4: "close your eyes",
    8: "jump",
  }

  result := []string{}

  for i, _ := range binaryCode {
    position := len(binaryCode)-i-1
    bit := int(binaryCode[position] - '0')

    if (bit > 0) {
      seekIndex := int(math.Pow(2, float64(i)))
      newCode := codes[seekIndex]
      if (newCode != "") {
        result = append(result, newCode)
      }
    }   
  }

  if (len(binaryCode) == 5) {
      newResult := []string{}
      backwardsResult := result
      for i, _ := range result {
        newResult = append(newResult, backwardsResult[len(backwardsResult) - i - 1])
      }
      result = newResult
    }

  return result
}
