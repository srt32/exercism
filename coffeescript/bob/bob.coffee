class Bob

  hey: (phrase) ->
    if @isSilence(phrase)
      'Fine. Be that way!'
    else if @isOrder(phrase)
      'Woah, chill out!'
    else if @isQuestion(phrase)
      'Sure.'
    else
      'Whatever.'

  isQuestion: (phrase) ->
    phrase.slice(-1) is '?'

  isOrder: (phrase) ->
    phrase.toUpperCase() is phrase

  isSilence: (phrase) ->
    !phrase or phrase.trim().length is 0

module.exports = Bob
