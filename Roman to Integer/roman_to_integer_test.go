package Roman_to_Integer

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		roman  string
		arabic int
	}{
		{"I", 1},
		{"II", 2},
		{"IV", 4},
		{"IX", 9},
		{"XIX", 19},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, tst := range tests {
		n := romanToInt(tst.roman)
		if n != tst.arabic {
			t.Errorf("%d is wrong answer, %s equal %d", n, tst.roman, tst.arabic)
		}
	}

}
