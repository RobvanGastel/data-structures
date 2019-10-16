package test

import (
	"reflect"
	"testing"
)

// AssertEqual check two variables if equal
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}

	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}
