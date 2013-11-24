var Words = function(input) {
  this.input = input;
  this.count = this.count_words();
};

Words.prototype.count_words = function() {
  split_words = this.clean_words(this.input.toLowerCase()).split(" ")
  var counts = {}
  for (var i = 0; i < split_words.length; i++) {
    counts[split_words[i]] = counts[split_words[i]] + 1 || 1;
  }
  return counts;
};

Words.prototype.clean_words = function(raw_words) {
  clean_words = raw_words.replace(/[^a-zA-Z0-9\s]/g, '').replace(/\s{2}/g, ' ');
  return clean_words;
};

module.exports = Words;
