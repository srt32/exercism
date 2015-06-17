package secret

import (
  "fmt"
  "strconv"
)

func Handshake(code int) []string {
  int64Code := int64(123)
  binaryCode := strconv.FormatInt(int64Code, 2)

  codes := map[int][]string{
    1: []string{"wink"},
    2: []string{"double blink"},
    4: []string{"close your eyes"},
    8: []string{"jump"},
  }

  //for 
  // loop over reverse of binaryCode and match each char to the codes map

  fmt.Println(codes[1])
  fmt.Println(binaryCode)

  return []string{"wink"}
}
