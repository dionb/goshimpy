goshimpy

[![GoDoc](https://godoc.org/github.com/dionb/goshimpy?status.svg)](https://godoc.org/github.com/dionb/goshimpy)

A go shim for python

When it's finished you will be able to call a function in python from go without needing to care how it happens.

To call a function in python call goshimpy.Call()

Parameters should be as follows: goshimpy.Call(funcName, importPath, &outputTarget, params…) 
