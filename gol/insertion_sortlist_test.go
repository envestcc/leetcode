package gol

import "testing"

type ExamInsertionSortList struct {
	input  string
	output string
}

func (e ExamInsertionSortList) Input() interface{} {
	return e.input
}

func (e ExamInsertionSortList) Output() interface{} {
	return e.output
}

func (e ExamInsertionSortList) Ac(answer interface{}) bool {
	answerSlice, ok := answer.(string)
	if !ok {
		return false
	}
	return answerSlice == e.Output()
}

func TestInsertionSortList(t *testing.T) {

	exams := []Exam{
		ExamInsertionSortList{
			input:  "4->2->1->3",
			output: "1->2->3->4",
		},
		ExamInsertionSortList{
			input:  "-1->5->3->4->0",
			output: "-1->0->3->4->5",
		},
	}

	ExamCheck(t, func(in interface{}) interface{} {
		head := ParseList(in.(string))
		return InsertionSortList(head).ToString()
	}, exams)

}
