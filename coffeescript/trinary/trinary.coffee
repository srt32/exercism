class Trinary

  constructor: (decimal) ->
    @reversedDec = @reverse(decimal)

  toDecimal: ->
    sum = 0
    unless @validInput(@reversedDec)
      sum += @convertDigit(digit, i) for digit, i in @reversedDec
    sum

  convertDigit: (digit, position) ->
    parseInt(digit, 10) * @threeToThe(position)

  validInput: (input) ->
    input.match(/[a-zA-Z]/)

  threeToThe: (power) ->
    Math.pow(3, power)

  reverse: (string) ->
    string.split("").reverse().join("")

module.exports = Trinary
