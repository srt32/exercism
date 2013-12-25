(ns bob)
(use 'clojure.string)
(declare silent?)
(declare yelling?)
(declare question?)

(defn response-for [phrase]
  (if (silent? phrase)
    "Fine. Be that way!"
  (if (yelling? phrase)
  "Woah, chill out!"
  (if (question? phrase)
    "Sure."
    "Whatever."
   )))
)

(defn- silent? [phrase]
 (or (nil? phrase) (blank? phrase))
)

(defn- yelling? [phrase]
  (= (upper-case phrase) phrase)
)

(defn- question? [phrase]
  (= \? (last phrase))
)
