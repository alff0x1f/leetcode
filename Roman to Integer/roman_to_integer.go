package Roman_to_Integer

func romanToInt(s string) int {
	costMap := map[rune]int{
		'I': 1, // can be minus
		'V': 5,
		'X': 10, // can be minus
		'L': 50,
		'C': 100, // can be minus
		'D': 500,
		'M': 1000,
	}
	minusPairs := []struct {
		minus rune
		main  rune
	}{
		{'I', 'V'}, // IV = 5-1 = 4
		{'I', 'X'}, // IX = 10-1 = 9
		{'X', 'L'}, // XL = 50-10 = 40
		{'X', 'C'}, // XC = 90-10 = 90
		{'C', 'D'}, // CD = 500-100 = 400
		{'C', 'M'}, // CM = 1000-100 = 900
	}
	prevValue := ' '
	sum := 0
	for _, letter := range s {
		sum += costMap[letter]
		for _, minusPair := range minusPairs {
			if prevValue == minusPair.minus && letter == minusPair.main {
				sum -= 2 * costMap[minusPair.minus]
			}
		}
		prevValue = letter
	}
	return sum
}
