package ex

import (
	"reflect"
	"testing"
)

func Pect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)",b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
