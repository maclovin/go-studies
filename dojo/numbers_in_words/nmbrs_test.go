package wrds

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "./nmbrs"
)

func TestWrds(t *testing.T) {
  inputValue := "745 \\$"
  outputValues := [2]string{"745.00 \\$", "seven hundred and fourty five dollars"}

  assert.Equal(t, inputValue, "745 \\$", "Testing basic input value.")
  assert.Equal(t, nmbrs.Wrds(inputValue), outputValues, "Testing decimal output value.")
  // assert.Equal(t, nmbrs.Wrds("1 \\$")[1], "one dollar", "Testing unit number conversion.")
}
