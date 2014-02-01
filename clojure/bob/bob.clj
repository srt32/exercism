(ns bob)
(use 'clojure.string)
(declare silent?)
(declare yelling?)
(declare question?)

(defn response-for [phrase]
  (cond
    (silent? phrase)   "Fine. Be that way!"
    (yelling? phrase)  "Woah, chill out!"
    (question? phrase) "Sure."
    :else "Whatever.")
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
