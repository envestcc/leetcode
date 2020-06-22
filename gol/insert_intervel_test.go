package gol

import "testing"

type Input struct {
	Intervals   [][]int
	NewInterval []int
}

type ExamInsertInterval struct {
	input  Input
	output [][]int
}

func (e ExamInsertInterval) Input() interface{} {
	return e.input
}

func (e ExamInsertInterval) Output() interface{} {
	return e.output
}

func (e ExamInsertInterval) Ac(answer interface{}) bool {
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

func TestInsertInterval(t *testing.T) {

	exams := []Exam{
		ExamInsertInterval{
			input: Input{
				Intervals:   [][]int{{1, 3}, {6, 9}},
				NewInterval: []int{2, 5},
			},
			output: [][]int{{1, 5}, {6, 9}},
		},
		ExamInsertInterval{
			input: Input{
				Intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				NewInterval: []int{4, 8},
			},
			output: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		ExamInsertInterval{
			input: Input{
				Intervals:   [][]int{},
				NewInterval: []int{5, 7},
			},
			output: [][]int{{5, 7}},
		},
		ExamInsertInterval{
			input: Input{
				Intervals:   [][]int{{1, 5}},
				NewInterval: []int{2, 3},
			},
			output: [][]int{{1, 5}},
		},
	}

	ExamCheck(t, func(in interface{}) interface{} {
		inp := in.(Input)
		return InsertInterval(inp.Intervals, inp.NewInterval)
	}, exams)

}
