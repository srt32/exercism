(ns dna)
(use 'clojure.string)
(declare convert)

(defn to-rna [letters]
  (if (some #{"X"} (map str letters))
    (throw (new AssertionError (str "AssertionError"))))
  (apply str (reduce concat (map convert (map str letters))))
)

(defn- convert [letter]
  (if (not= letter "T")
    (str letter)
    "U")
)
