package romanparse

import (
  "strings"
  "errors"
)

// Converts roman strings to int
func ToInteger(value string) (int, error) {
  value = strings.ToUpper(value)
  value = strings.TrimSpace(value)
  max := len(value)

  if max == 0 {
    return 0, errors.New("empty value")
  }

  roman := map[string]int{
    "I": 1, "V": 5, "X": 10,
    "L": 50, "C": 100, "D": 500, "M": 1000,
  }

  output := 0
  rpkey := 0
  last := max - 1
  lastkey := string(value[last])
  lnumber, _ := roman[lastkey]
  
  for i := last; i >= 0; i-- {
    key := string(value[i])
    number, found := roman[key]
    if !found {
      return 0, errors.New("invalid value")
    }

    if number >= lnumber {      
      output += number
    } else {
      output += -number
    }

    if i < last {
      if lastkey == key {
        rpkey++
      } else {
        rpkey = 0
      }
      if rpkey >= 3 {
        return 0, errors.New("invalid value")
      }

      lastkey = string(value[i])
      lnumber, _ = roman[lastkey]
    }
  }

  return output, nil
}
