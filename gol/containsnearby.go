package gol

import "math"

func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	buckets := map[int]int{}
	w := t + 1
	for i := range nums {
		m := getBucketID(nums[i], w)
		if _, ok := buckets[m]; ok {
			return true
		}
		buckets[m] = nums[i]

		if v, ok := buckets[m-1]; ok && int(math.Abs(float64(v-nums[i]))) < w {
			return true
		}
		if v, ok := buckets[m+1]; ok && int(math.Abs(float64(v-nums[i]))) < w {
			return true
		}

		if i >= k {
			delete(buckets, getBucketID(nums[i-k], w))
		}
	}
	return false
}

func getBucketID(a, w int) int {
	if a >= 0 {
		return a / w
	} else {
		return (a+1)/w - 1
	}
}
