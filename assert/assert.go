package assert

import (
	"reflect"
)

func Equal[T, U any](l T, r U, shouldBe bool) bool {
	return reflect.DeepEqual(l, r) == shouldBe
}
