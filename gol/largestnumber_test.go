package gol

import "testing"

type ExamLargestNumber struct {
	input  []int
	output string
}

func (e ExamLargestNumber) Input() interface{} {
	return e.input
}

func (e ExamLargestNumber) Output() interface{} {
	return e.output
}

func (e ExamLargestNumber) Ac(answer interface{}) bool {
	answerSlice, ok := answer.(string)
	if !ok {
		return false
	}
	return answerSlice == e.Output()
}

func TestLargestNumber(t *testing.T) {

	exams := []Exam{
		ExamLargestNumber{
			input:  []int{10, 2},
			output: "210",
		},
		ExamLargestNumber{
			input:  []int{3, 30, 34, 5, 9},
			output: "9534330",
		},
		ExamLargestNumber{
			input:  []int{0, 0},
			output: "0",
		},
	}

	ExamCheck(t, func(in interface{}) interface{} {
		return LargestNumber(in.([]int))
	}, exams)

}
