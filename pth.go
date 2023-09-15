package pth

import (
	"fmt"
	"github.com/RealPatricia/pth/assert"
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
			if !assert.Equal(result, test.Expecting, test.Success) {
				t.Errorf(err(test.Expecting, result))
			}
		})
	}
}

func SanityTest(t *testing.T) {
	if 2+2 != 4 {
		t.Errorf("Sorry folks, math's closed, moose out front should have told ya")
	}

	type in struct {
		l, r int
	}

	tests := []Test[in, int]{
		{
			Name:      "five",
			Input:     in{2, 3},
			Expecting: 5,
			Success:   true,
		},
		{
			Name:      "seven",
			Input:     in{2, 3},
			Expecting: 7,
			Success:   false,
		},
	}

	f := func(val in) int {
		return val.l + val.r
	}

	Run(tests, f, t)
}
