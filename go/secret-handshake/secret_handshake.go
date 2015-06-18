package secret

import (
  "fmt"
  "math"
  "strconv"
)

func Handshake(code int) []string {
  binaryCode := strconv.FormatInt(int64(code), 2)

  codes := map[int]string{
    1: "wink",
    2: "double blink",
    4: "close your eyes",
    8: "jump",
  }

  fmt.Println("binaryCode is: ", binaryCode)

  result := []string{}

  for i, _ := range binaryCode {
    index := int(binaryCode[len(binaryCode)-i-1] - '0')

    if (index > 0) {
      newCode := codes[int(math.Pow(2, float64(len(binaryCode)-i-1)))]
      result = append(result, newCode)
    }  
  }
  // loop over reverse of binaryCode and match each char to the codes map

  return result
}
