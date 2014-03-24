class Bob

  hey: (phrase) ->
    if @.isSilence(phrase)
      'Fine. Be that way!'
    else if @.isOrder(phrase)
      'Woah, chill out!'
    else if @.isQuestion(phrase)
      'Sure.'
    else
      'Whatever.'

  isQuestion: (phrase) ->
    phrase.slice(-1) == '?'

  isOrder: (phrase) ->
    phrase.toUpperCase() == phrase

  isSilence: (phrase) ->
    !phrase || phrase.trim().length == 0

module.exports = Bob
