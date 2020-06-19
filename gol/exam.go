package gol

import (
	"reflect"
	"testing"
)

type Exam interface {
	Input() interface{}
	Output() interface{}
	Ac(interface{}) bool
}

func ExamCheck(t *testing.T, method func(interface{}) interface{}, exams []Exam) {
	for _, e := range exams {
		output := method(e.Input())
		if !e.Ac(output) {
			t.Errorf("function [%+v] check fail.\nInput=%+v\nOutput=%+v\nYours=%+v", reflect.TypeOf(method).Name(), e.Input(), e.Output(), output)
		} else {
			// t.Logf("function [%+v] check success.\nInput=%+v\nOutput=%+v\nYours=%+v", reflect.TypeOf(method).Name(), e.Input(), e.Output(), output)
		}
	}
}
