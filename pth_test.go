package pth_test

import (
	"github.com/RealPatricia/pth"
	"github.com/RealPatricia/pth/assert"
	"testing"
)

func TestSanity(t *testing.T) {
	pth.SanityTest(t)

	if !assert.Equal(2+2, 4, true) {
		t.Errorf("Better luck next time chucklefuck")
	}
}
