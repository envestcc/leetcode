package gol

import (
	"fmt"
	"sort"
	"strings"
)

func LargestNumber(nums []int) string {

	lessThan := func(i, j int) bool {
		// fmt.Printf("%+v ? %+v\n", nums[i], nums[j])
		is := partsint(nums[i])
		js := partsint(nums[j])

		ii, ji := len(is)-1, len(js)-1
		for {
			for ii >= 0 && ji >= 0 {
				if is[ii] > js[ji] {
					// fmt.Printf("%+v > %+v\n", is[ii], js[ji])
					return false
				} else if is[ii] < js[ji] {
					// fmt.Printf("%+v < %+v\n", is[ii], js[ji])
					return true
				}
				// fmt.Printf("%+v = %+v\n", is[ii], js[ji])
				ii--
				ji--
			}
			if ii >= 0 {
				ji = len(js) - 1
			} else if ji >= 0 {
				ii = len(is) - 1
			} else {
				return true
			}
		}

	}
	largeThan := func(i, j int) bool {
		return !lessThan(i, j)
	}
	sort.Slice(nums, largeThan)

	var result strings.Builder
	noZero := false
	zeroCnt := 0
	for i := range nums {
		if !noZero && nums[i] == 0 {
			if zeroCnt == 0 {
				zeroCnt++
			} else {
				continue
			}

		}
		if nums[i] != 0 {
			noZero = true
		}

		result.WriteString(fmt.Sprintf("%d", nums[i]))
	}
	return result.String()
}

func partsint(n int) []int {
	result := []int{}
	for {
		result = append(result, n%10)
		n /= 10
		if n == 0 {
			break
		}
	}
	return result
}
