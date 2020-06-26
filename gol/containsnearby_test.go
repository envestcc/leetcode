package gol

import (
	"testing"
)

type InputContainsNearby struct {
	Nums []int
	K    int
	T    int
}

type ExamContainsNearby struct {
	input  InputContainsNearby
	output bool
}

func (e ExamContainsNearby) Input() interface{} {
	return e.input
}

func (e ExamContainsNearby) Output() interface{} {
	return e.output
}

func (e ExamContainsNearby) Ac(answer interface{}) bool {
	answerSlice, ok := answer.(bool)
	if !ok {
		return false
	}
	return answerSlice == e.output
}

func TestContainsNearby(t *testing.T) {

	exams := []Exam{
		ExamContainsNearby{
			input: InputContainsNearby{
				Nums: []int{1, 2, 3, 1},
				K:    3,
				T:    0,
			},
			output: true,
		},
		ExamContainsNearby{
			input: InputContainsNearby{
				Nums: []int{1, 0, 1, 1},
				K:    1,
				T:    2,
			},
			output: true,
		},
		ExamContainsNearby{
			input: InputContainsNearby{
				Nums: []int{1, 5, 9, 1, 5, 9},
				K:    2,
				T:    3,
			},
			output: false,
		},
	}

	ExamCheck(t, func(in interface{}) interface{} {
		input := in.(InputContainsNearby)
		return ContainsNearbyAlmostDuplicate(input.Nums, input.K, input.T)
	}, exams)

}
