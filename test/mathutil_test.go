package sum

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "./mathutil"
)

func TestSum(t *testing.T) {
  testRes := mathutil.Sum(1, 2)
  assert.Equal(t, testRes, 3, "True is true! In other worlds: 1 + 2 = 3")
}
