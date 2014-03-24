class Binary
  constructor: (@binary) ->
    @reversedDigits = binary.split("").reverse().join("")

  toDecimal: ->
    sum = 0
    unless @reversedDigits.match(/[a-zA-Z]/)
      for digit, index in @reversedDigits
        sum += parseInt(digit, 10)*Math.pow(2, index)
    sum

module.exports = Binary
