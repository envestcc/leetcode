package gol

import (
	"testing"
)

type InputFindKthLargest struct {
	Nums []int
	K    int
}

type ExamFindKthLargest struct {
	input  InputFindKthLargest
	output int
}

func (e ExamFindKthLargest) Input() interface{} {
	return e.input
}

func (e ExamFindKthLargest) Output() interface{} {
	return e.output
}

func (e ExamFindKthLargest) Ac(answer interface{}) bool {
	answerSlice, ok := answer.(int)
	if !ok {
		return false
	}
	return answerSlice == e.output
}

func TestFindKthLargest(t *testing.T) {

	exams := []Exam{
		ExamFindKthLargest{
			input: InputFindKthLargest{
				Nums: []int{3, 2, 1, 5, 6, 4},
				K:    2,
			},
			output: 5,
		},
		ExamFindKthLargest{
			input: InputFindKthLargest{
				Nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				K:    4,
			},
			output: 4,
		},
		ExamFindKthLargest{
			input: InputFindKthLargest{
				Nums: []int{2, 1},
				K:    2,
			},
			output: 1,
		},
	}

	ExamCheck(t, func(in interface{}) interface{} {
		input := in.(InputFindKthLargest)
		return FindKthLargest(input.Nums, input.K)
	}, exams)

}
