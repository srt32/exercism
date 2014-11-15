module Bob (responseFor) where

import Data.Char

responseFor :: String -> String
responseFor statement =
   if isSilence statement
     then "Fine. Be that way!"
     else
      if isYelling statement
        then "Woah, chill out!"
        else
          if isQuestion statement
            then "Sure."
            else "Whatever."

isSilence :: String -> Bool
isSilence statement =
  (null statement) || (last statement) == ' '

isYelling :: String -> Bool
isYelling statement =
// if it contains any letters AND
  (map toUpper statement) == statement

isQuestion :: String -> Bool
isQuestion statement =
  (last statement) == '?'
