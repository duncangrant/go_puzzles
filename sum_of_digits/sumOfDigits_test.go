package sumOfDigits

import (
  "testing"
  "strconv"
  "math"
)

var expectedResults = map[string]int{
  "1":-1,
  "11":20,
  "12":21,
  "30":-1,
}

func findNextNumberWithMatchingDigits(digits string) int {
  i,_ := strconv.Atoi(digits)
  i = i + 9
  if int(math.Floor(math.Log(float64(i)))) > len(digits) {
    return -1
  }
  return i
}

func Test(t *testing.T) {
  for item := range expectedResults {
    if result := findNextNumberWithMatchingDigits(item); result != expectedResults[item] {
      t.Errorf("Fail %s expected %d got %d", item, expectedResults[item], result)
    }
    t.Log("Pass %s got %d", item, expectedResults[item])
  }
}
