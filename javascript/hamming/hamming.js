var Hamming = {};

Hamming.compute = function(baseStrand, copyStrand){
  var distance = 0;
  for (var i = 0; i < baseStrand.length; i++) {
    if (notEqual(baseStrand[i], copyStrand[i])) {
      distance++;
    }
  }
  return distance;
};

var notEqual = function(base, copy) {
  return copy && (base != copy);
};

module.exports = Hamming;
