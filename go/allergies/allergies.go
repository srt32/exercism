package allergies

import(
  "strconv"
)

func Allergies(score int) (allergyList []string) {
  binary := strconv.FormatInt(int64(score), 2)
  for i, _ := range binary {
    allergyCode :=
    allergyList = append(allergyList, foodForCode(allergyCode))
  }
  return allergyList
}

func AllergicTo(i int, allergen string) bool {
  return false
}

func foodForCode(code int) (food string) {
  return map[int]string{
    0: "",
    1: "eggs",
  }[code]
}

