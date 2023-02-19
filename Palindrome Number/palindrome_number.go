package Palindrome_Number

func isPalindrome(x int) bool {
	if x < 0 {
		// negative numbers (like -5) can't be a palindrome because number can't end
		// with minus sign
		return false
	}
	digits := splitIntToDigits(x)
	for i := 0; i < len(digits); i++ {
		if digits[i] != digits[len(digits)-i-1] {
			return false
		}
	}
	return true
}

func splitIntToDigits(x int) []int {
	// split number to single digits
	reverseDigits := make([]int, 0)
	for x > 0 {
		reverseDigits = append(reverseDigits, x%10)
		x = x / 10
	}
	countDigits := len(reverseDigits)
	digits := make([]int, countDigits)
	for idx := range reverseDigits {
		digits[idx] = reverseDigits[countDigits-idx-1]
	}
	return digits
}
