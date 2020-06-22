package gol

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	// sort
	left, right := 0, len(intervals)-1
	index := 0
	for {
		// fmt.Printf("find in [%+v, %+v]\n", left, right)
		if right < 0 {
			break
		}
		if left == right {
			if intervals[right][0] >= newInterval[0] {
				index = right
			} else {
				index = right + 1
			}

			break
		}
		mid := (left + right) / 2
		if intervals[mid][0] == newInterval[0] {
			index = mid
			break
		}
		if intervals[mid][0] > newInterval[0] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	intervals = append(intervals, []int{0, 0})
	for i := len(intervals) - 2; i >= index; i-- {
		intervals[i+1] = intervals[i]
	}
	intervals[index] = newInterval
	// fmt.Printf("index = %+v\n%+v\n", index, intervals)

	// merge
	copyIntervals := intervals
	result := make([][]int, 0)
	for i := range copyIntervals {
		if len(result) == 0 {
			result = append(result, []int(copyIntervals[i]))
		} else {
			if copyIntervals[i][0] <= result[len(result)-1][1] { // merge
				if copyIntervals[i][1] > result[len(result)-1][1] {
					result[len(result)-1][1] = copyIntervals[i][1]
				}
			} else { // add
				result = append(result, []int(copyIntervals[i]))
			}
		}
	}

	return result
}
