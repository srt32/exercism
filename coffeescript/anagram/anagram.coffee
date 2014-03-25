class Anagram

  constructor: (targetWord) ->
    @targetWord = targetWord.toLowerCase()

  match: (words) ->
    matches = []
    for word in words
      word = word.toLowerCase()
      matches.push(word) if @isAnagram(word)
    matches

  isAnagram: (secondWord) ->
    secondWord != @targetWord && @sort(@targetWord) == @sort(secondWord)

  sort: (word) ->
    word.toLowerCase().split('').sort().join('')

module.exports = Anagram
