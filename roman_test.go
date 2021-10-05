package romanparse

import (
  "testing"
)

type TestCase struct {
  input string
  output int
  error bool
}

func TestToInteger(t *testing.T) {
  cases := []TestCase{
    {"", 0, true},
    {"XXXXV", 0, true},
    {"ABC", 0, true},
    {" ABC ", 0, true},
    {"XVAC", 0, true},
    {"123", 0, true},
    {"---", 0, true},
    {"0", 0, false},
    {"XXX", 30, false},
    {"XV", 15, false},
    {"XIV", 14, false},
    {"XXVI", 26, false},
    {"MMXXI", 2021, false},
    {"mmccxxii", 2222, false},
  }

  for _, test := range cases {
    result, err := ToInteger(test.input)
    if test.error {
      if err != nil {
        t.Logf("ToInteger('%v'): PASSED, expected error and got error '%v'", test.input, err)
      } else {
        t.Errorf("ToInteger('%v'): FAILED, expected error but got value '%v'", test.input, result)
      }
    } else {
      if result == test.output {
        t.Logf("ToInteger('%v'): PASSED, expected '%v' and got value '%v'", test.input, test.output, result)
      } else {
        t.Errorf("ToInteger('%v'): FAILED, expected '%v' but got value '%v'", test.input, test.output, result)
      }
    }
  }
}
