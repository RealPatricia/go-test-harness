package pth

import (
	"fmt"
	"reflect"
	"testing"
)

type Test[T, U any] struct {
	Name      string
	Input     T
	Expecting U
	Success   bool
}

func err[T any](expected, result T) string {
	return fmt.Sprintf("\nExpecting:\t%v\nGot:\t%v", expected, result)
}

func Run[T, U any](tests []Test[T, U], f func(in T) U, t *testing.T) {
	for _, test := range tests {
		name := test.Name
		t.Run(name, func(t *testing.T) {
			result := f(test.Input)
			if test.Success != reflect.DeepEqual(test.Expecting, result) {
				t.Errorf(err(test.Expecting, result))
			}
		})
	}
}
