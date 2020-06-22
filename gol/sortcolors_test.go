package gol

import "testing"

type ExamSortColors struct {
	input  []int
	output []int
}

func (e ExamSortColors) Input() interface{} {
	return e.input
}

func (e ExamSortColors) Output() interface{} {
	return e.output
}

func (e ExamSortColors) Ac(answer interface{}) bool {
	answerSlice, ok := answer.([]int)
	if !ok {
		return false
	}
	if len(e.output) != len(answerSlice) {
		return false
	}
	for i := range e.output {
		if e.output[i] != answerSlice[i] {
			return false
		}
	}
	return true
}

func TestSortColors(t *testing.T) {

	exams := []Exam{
		ExamSortColors{
			input:  []int{2, 0, 2, 1, 1, 0},
			output: []int{0, 0, 1, 1, 2, 2},
		},
	}

	ExamCheck(t, func(in interface{}) interface{} {
		return SortColors(in.([]int))
	}, exams)

}
