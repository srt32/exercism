var Anagram = function(input_string) {
  this.string = input_string;
};

Anagram.prototype.match = function(list) {
  var matching_words = []
  for(var i = 0; i < list.length; i++) {
    var local_word = list[i];
    if (this.isAnagramFor(local_word)){
      matching_words.push(local_word);
    };
  };
  return matching_words;
};

Anagram.prototype.isAnagramFor = function(local_word) {
  if (this.count(local_word) === this.count(this.string) && local_word !== this.string){
    return true;
  };
};

Anagram.prototype.count = function(input) {
  var word_count = 0
    for(var j = 0; j < input.length; j++){
      word_count += input.charCodeAt(j);
    };
  return word_count;
};

module.exports = Anagram;
