package wc

import (
  "strings"
  "unicode"
)

type Histogram map[string]int

func WordCount(input string) Histogram {
  h := make(Histogram)
  rawWords := strings.Fields(input)

  for _, word := range rawWords {
    cleanWord := []rune{}
    if word != ":" {
      for _, c := range word {
        if validCharacter(c) {
          cleanWord = append(cleanWord, c)
        }
      }
      h[makeLowerWord(cleanWord)] += 1
    }
  }
  return h
}

func validCharacter(c rune) bool {
  return unicode.IsLetter(c) || unicode.IsNumber(c)
}

func makeLowerWord(word []rune) string {
  return strings.ToLower(string(word))
}
