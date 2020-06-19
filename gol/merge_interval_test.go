package gol

import "testing"

type ExamMergeInterval struct {
	input  [][]int
	output [][]int
}

func (e ExamMergeInterval) Input() interface{} {
	return e.input
}

func (e ExamMergeInterval) Output() interface{} {
	return e.output
}

func (e ExamMergeInterval) Ac(answer interface{}) bool {
	answerSlice, ok := answer.([][]int)
	if !ok {
		return false
	}
	if len(e.output) != len(answerSlice) {
		return false
	}
	for i := range e.output {
		if len(e.output[i]) != len(answerSlice[i]) {
			return false
		}
		for j := range e.output[i] {
			if e.output[i][j] != answerSlice[i][j] {
				return false
			}
		}
	}
	return true
}

func TestMergeInterval(t *testing.T) {

	exams := []Exam{
		ExamMergeInterval{
			input:  [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			output: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		ExamMergeInterval{
			input:  [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			output: [][]int{{1, 10}},
		},
	}

	ExamCheck(t, func(in interface{}) interface{} {
		return MergeInterval(in.([][]int))
	}, exams)

}
