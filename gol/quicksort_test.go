package gol

import "testing"

func TestQuickSort(t *testing.T) {
	// data1 := []int{5, 9, 10, 1, 3, 7, 12, 90, 12, 23, 56, 89, 120, 8, 12, 23, 45, 90}
	data := []int{0, 2, 4, 0, 5, 4, 0}
	Sort(data, func(i, j int) bool { return data[i] < data[j] })
	t.Log(data)
}
