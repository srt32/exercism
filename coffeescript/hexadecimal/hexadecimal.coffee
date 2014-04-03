class Hexadecimal

  constructor: (hex) ->
    @reverseHex = @_reverse(hex)

  toDecimal: ->
    return 0 unless @_validHex()
    result = 0
    for digit, index in @reverseHex
      result += @_toDec(digit) * Math.pow(16, index)
    result

  _toDec: (digit) ->
    {
      'a': 10,
      'b': 11,
      'c': 12,
      'd': 13,
      'e': 14,
      'f': 15
    }[digit] or digit

  _validHex: ->
    @reverseHex.match(/^[a-f0-9]*$/)

  _reverse: (hex) ->
    hex.split('').reverse().join('')

module.exports = Hexadecimal
