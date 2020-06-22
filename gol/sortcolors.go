package gol

func SortColors(nums []int) []int {
	kv := make(map[int]int)
	for _, v := range nums {
		if count, ok := kv[v]; ok {
			kv[v] = count + 1
		} else {
			kv[v] = 1
		}
	}
	index := 0
	if count, ok := kv[0]; ok {
		for i := 0; i < count; i++ {
			nums[i+index] = 0
		}
		index += count
	}
	if count, ok := kv[1]; ok {
		for i := 0; i < count; i++ {
			nums[i+index] = 1
		}
		index += count
	}
	if count, ok := kv[2]; ok {
		for i := 0; i < count; i++ {
			nums[i+index] = 2
		}
		index += count
	}
	return nums
}
