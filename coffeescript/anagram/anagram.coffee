class Anagram

  constructor: (targetWord) ->
    @targetWord = targetWord.toLowerCase()

  match: (words) ->
    matches = []
    for word in words
      word = word.toLowerCase()
      unless word == @targetWord
        sortedWord = word.toLowerCase().split('').sort().join('')
        if @targetWord.split('').sort().join('') == sortedWord
          matches.push(word)
    matches

module.exports = Anagram
