package gol

import "testing"

type ExamSortList struct {
	input  string
	output string
}

func (e ExamSortList) Input() interface{} {
	return e.input
}

func (e ExamSortList) Output() interface{} {
	return e.output
}

func (e ExamSortList) Ac(answer interface{}) bool {
	answerSlice, ok := answer.(string)
	if !ok {
		return false
	}
	return answerSlice == e.Output()
}

func TestSortList(t *testing.T) {

	exams := []Exam{
		ExamSortList{
			input:  "4->2->1->3",
			output: "1->2->3->4",
		},
		ExamSortList{
			input:  "-1->5->3->4->0",
			output: "-1->0->3->4->5",
		},
	}

	ExamCheck(t, func(in interface{}) interface{} {
		head := ParseList(in.(string))
		return SortList(head).ToString()
	}, exams)

}
