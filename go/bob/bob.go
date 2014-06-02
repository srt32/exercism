package bob

import (
  "strings"
)

func Hey(sentence string) string {
  if isShouting(sentence) {
    return "Woah, chill out!"
  } else if isQuestioning(sentence) {
    return "Sure."
  } else {
    return "Whatever."
  }
}

func isShouting(s string) bool {
  return strings.ToUpper(s) == s && !onlyNumbers(s)
}

func isQuestioning(s string) bool {
  return strings.HasSuffix(s, "?")
}

func onlyNumbers(s string) bool {

  return false
}
