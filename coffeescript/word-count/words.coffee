class Words
  constructor: (sentence) ->
    sentence = sentence.toLowerCase()
    words = sentence.match(/[a-z0-9]+/g)
    @count = {}

    for word in words
      @count[word] = 0 unless @count[word]?
      @count[word] += 1

  module.exports = @
