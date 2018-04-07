package goshimpy_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/dionb/goshimpy"
)

func TestCall(t *testing.T) {
	var output string
	err := goshimpy.Call("test_print", "test", &output, "HELLO WORLD")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(output)
	if output != "HELLO WORLD" {
		t.Fail()
	}
}

func TestCallMultipleArgs(t *testing.T) {
	var output string
	err := goshimpy.Call("test_print2", "test", &output, "HELLO", "WORLD")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(output)
	if output != "HELLO WORLD" {
		t.Fail()
	}
}

func TestCallOtherDir(t *testing.T) {
	var output string
	err := goshimpy.Call("foo", "test_dir.child_test", &output)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(output)
	if output != "bar" {
		t.Fail()
	}
}
