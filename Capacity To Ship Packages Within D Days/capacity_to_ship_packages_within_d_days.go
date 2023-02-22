package Capacity_To_Ship_Packages_Within_D_Days

func shipDuration(weight *[]int, capacity int) (days int, maxLoad int) {
	days = 1
	maxLoad = 0
	load := 0
	for _, w := range *weight {
		load += w
		if load > capacity {
			days++
			load = w
		}
		if load > maxLoad {
			maxLoad = load
		}
	}
	return days, maxLoad
}

func shipWithinDays(weights []int, days int) int {
	capacityMinBound := 0
	capacityMaxBound := 0
	for _, w := range weights {
		if w > capacityMinBound {
			// cargo can't be smaller than bigger package
			capacityMinBound = w
		}
		// can ship by 1 day
		capacityMaxBound += w
	}
	var maxLoad int
	for capacityMinBound < capacityMaxBound {
		middleCapacity := (capacityMaxBound-capacityMinBound)/2 + capacityMinBound
		if middleCapacity == capacityMaxBound {
			// infinity loop protection
			middleCapacity = capacityMinBound
		}
		var shipDays int
		shipDays, maxLoad = shipDuration(&weights, middleCapacity)
		if shipDays == days {
			capacityMaxBound = maxLoad
		}
		if shipDays > days {
			// take longer than need, capacity not fit
			capacityMinBound = middleCapacity + 1
		}
		if shipDays < days {
			// shipDays can be increased with non-optimal cargo method, so
			// middleCapacity can be answer if middleCapacity+1 doesn't fit
			capacityMaxBound = middleCapacity
		}
	}
	return capacityMaxBound
}
