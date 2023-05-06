package Average_Salary_Excluding_the_Minimum_and_Maximum_Salary

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	salary := []int{6000, 5000, 4000, 3000, 2000, 1000}
	expected := 3500.00000
	result := average(salary)
	if math.Abs(result-expected) > 1e-5 {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestAverageWithNegativeValues(t *testing.T) {
	salary := []int{-6000, -5000, 4000, 3000, 2000, 1000}
	expected := 250.00000
	result := average(salary)
	if math.Abs(result-expected) > 1e-5 {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestAverageWithOneValue(t *testing.T) {
	salary := []int{5000}
	expected := 5000.00000
	result := average(salary)
	if math.Abs(result-expected) > 1e-5 {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
