package allergies

import (
  "math"
  "strconv"
)

func Allergies(code int) (allergens []string) {
  binary := strconv.FormatInt(int64(code), 2)

  var allergenCodes []float64

  for i, _ := range binary {
    index := len(binary) - i - 1
    value := string(binary[i])

    floatValue, _ := strconv.ParseFloat(value, 64)

    allergenCode := math.Pow(2, float64(index)) * floatValue
    allergenCodes = append(allergenCodes, allergenCode)
  }

  for _, value := range allergenCodes {
    allergen := allergenForCode(int(value))
    if len(allergen) > 0 {
      allergens = append(allergens, allergen)
    }
  }

  length := len(allergens)
  sortedAllergens := make([]string, length)
  for i, _ := range allergens {
    sortedAllergens[i] = allergens[length-1-i]
  }

  return sortedAllergens
}

func allergenForCode(i int) string {
  return map[int]string{
    1:   "eggs",
    2:   "peanuts",
    4:   "shellfish",
    8:   "strawberries",
    16:  "tomatoes",
    32:  "chocolate",
    64:  "pollen",
    128: "cats",
  }[i]
}

func AllergicTo(i int, allergen string) bool {
  allergies := Allergies(i)

  for _, allergy := range allergies {
    if allergy == allergen {
      return true
    }
  }
  return false
}
