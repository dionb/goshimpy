package goshimpy

import (
	"fmt"
	"os/exec"
)

// var SafeMode bool

//Call a python function using the python environment
func Call(name, path string, result interface{}, params ...interface{}) error {

	toPrint := params[0]

	cmd := exec.Command("python", "-c", fmt.Sprintf("import %s; print %s.%s(%s)", path, path, name, toPrint))
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(out))

	return nil

}
