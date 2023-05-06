package Average_Salary_Excluding_the_Minimum_and_Maximum_Salary

func average(salary []int) float64 {
	maxVal, minVal := salary[0], salary[0]
	sum := 0
	for _, sal := range salary {
		sum += sal
		if maxVal < sal {
			maxVal = sal
		}
		if minVal > sal {
			minVal = sal
		}
	}
	return float64(sum-minVal-maxVal) / float64(len(salary)-2)
}
