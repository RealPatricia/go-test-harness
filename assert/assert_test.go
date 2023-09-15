package assert_test

import (
	"github.com/RealPatricia/pth/assert"
	"testing"
)

func TestAsset(t *testing.T) {
	if !assert.Equal(4, 4, true) || !assert.Equal(3, 4, false) {
		t.Errorf("Either your code is wrong, or math is broken")
	}

}
