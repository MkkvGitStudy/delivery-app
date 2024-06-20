package util

import "math"

const FilePath = "offer_codes.json"

func Min(nums ...float64) (float64, int) {
	if len(nums) == 0 {
		return -1, -1
	}
	min := nums[0]
	ind := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			ind = i
		}
	}
	return min, ind
}

func Round(val float64) float64 {
	return math.Round(val*100) / 100
}
