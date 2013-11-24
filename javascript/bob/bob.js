var Bob = function() {
  this.hey = function (phrase) {
    if (silence(phrase)) {
      return 'Fine. Be that way!'
    }
    else if (yelling(phrase)) {
      return 'Woah, chill out!'
    }
    else if (question(phrase)) {
      return 'Sure.'
    }
    else
      return 'Whatever.'
  }
  silence = function(phrase) {
    return phrase.match(/^\s*$/)
  }
  question = function(phrase) {
    return phrase.charAt(phrase.length-1) == '?'
  }
  yelling = function(phrase) {
    return phrase.toUpperCase() == phrase
  }
};
module.exports = Bob;
