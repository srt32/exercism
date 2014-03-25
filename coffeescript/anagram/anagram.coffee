class Anagram

  constructor: (targetWord) ->
    @targetWord = targetWord.toLowerCase()

  match: (words) ->
    matches = []
    for word in words
      word = word.toLowerCase()
      matches.push(word) if @isAnagram(@targetWord, word)
    matches

  isAnagram: (firstWord, secondWord) ->
    secondWord != @targetWord && @sort(firstWord) == @sort(secondWord)

  sort: (word) ->
    word.toLowerCase().split('').sort().join('')

module.exports = Anagram
