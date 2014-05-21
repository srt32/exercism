package scrabble_score

import (
  "strings"
)

func Score(s string) int {
  myWord := scrabble{s}
  return myWord.scoreWord()
}

type scrabble struct {
  word string
}

func (s *scrabble) scoreWord() (total int) {
  for _, letter := range s.letters() {
    total += s.scoreLetter(letter)
  }
  return total
}

func (s *scrabble) letters() []string {
  return strings.Split(s.word, "")
}

func (s *scrabble) scoreLetter(letter string) (score int) {
  return map[string]int{
    "a": 1,
    "b": 3,
    "c": 3,
    "d": 2,
    "e": 1,
    "f": 4,
    "g": 2,
    "h": 4,
    "i": 1,
    "j": 8,
    "k": 5,
    "l": 1,
    "m": 3,
    "n": 1,
    "o": 1,
    "p": 3,
    "q": 10,
    "r": 1,
    "s": 1,
    "t": 1,
    "u": 1,
    "v": 4,
    "w": 4,
    "x": 8,
    "y": 4,
    "z": 10,
  }[strings.ToLower(letter)]
}
