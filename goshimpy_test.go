package goshimpy_test

import (
	"testing"

	"github.com/dionb/goshimpy"
)

func TestCall(t *testing.T) {
	goshimpy.Call("test_print", "test", "\"HELLO WORLD\"")
	t.Fail()
}
