package gol

import "sort"

/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

链接：https://leetcode-cn.com/problems/merge-intervals
*/

func MergeInterval(intervals [][]int) [][]int {

	copyIntervals := make([][]int, len(intervals))
	for i := range intervals {
		copyIntervals[i] = make([]int, len(intervals[i]))
		copy(copyIntervals[i], intervals[i])
	}

	sort.Slice(copyIntervals, func(i, j int) bool {
		return copyIntervals[i][0] < copyIntervals[j][0]
	})

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
