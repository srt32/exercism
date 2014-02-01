var Beer = function() {
};

Beer.verse = function(round){
  var singular = (round === 1);
  if (round === 0){
    return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
  };
  result = [round, " bottle", (singular ? "" : "s"), " of beer on the wall, ",
            round, " bottle", (singular ? "" : "s"), " of beer.\nTake ",
            (singular ? 'it' : 'one'),
            " down and pass it around, ",
            (singular ? 'no more' : round - 1),
            " bottle", (singular ? "" : "s"), " of beer on the wall.\n"].join("");
  return result;
};

Beer.sing = function(start, end){
  end = end || 0;
  var song = "";
  var i;
  for(i = start; i >= end; i--){
    //song = [song, Beer.verse(i)];
    song = song + Beer.verse(i) + "\n";
  };
  return song.slice(0, -1);
};

module.exports = Beer;
