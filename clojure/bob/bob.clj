(ns bob)
(use 'clojure.string)
(def response-for
  (fn [phrase]
    (if (or (nil? phrase) (blank? phrase))
      "Fine. Be that way!"
    (if (= (upper-case phrase) phrase)
      "Woah, chill out!"
    (if (= \? (last phrase))
      "Sure."
      "Whatever."
    )))
  )
)
