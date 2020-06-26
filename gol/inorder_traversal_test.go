package gol

import (
	"testing"
)

type ExamInOrderTraversal struct {
	input  string
	output []int
}

func (e ExamInOrderTraversal) Input() interface{} {
	return e.input
}

func (e ExamInOrderTraversal) Output() interface{} {
	return e.output
}

func (e ExamInOrderTraversal) Ac(answer interface{}) bool {
	answerSlice, ok := answer.([]int)
	if !ok {
		return false
	}
	if len(answerSlice) != len(e.output) {
		return false
	}
	for i := range answerSlice {
		if answerSlice[i] != e.output[i] {
			return false
		}
	}
	return true
}

func TestInOrderTraversal(t *testing.T) {

	exams := []Exam{
		ExamInOrderTraversal{
			input:  "1,null,2,3",
			output: []int{1, 3, 2},
		},
		ExamInOrderTraversal{
			input:  "2,3,1",
			output: []int{1, 3, 2},
		},
	}

	ExamCheck(t, func(in interface{}) interface{} {
		root := ParseTreeFromString(in.(string))
		return InOrderTraversal(root)
	}, exams)

}
