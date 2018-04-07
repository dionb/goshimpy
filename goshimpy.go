package goshimpy

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// var SafeMode bool

//Call a python function using the python environment
func Call(name, path string, result interface{}, params ...interface{}) error {

	toPrint := params[0]

	cmd := exec.Command("python", "-c", fmt.Sprintf("import json,%s; print(json.dumps(%s.%s(%s)))", path, path, name, toPrint))
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(out))
		return err
	}
	fmt.Println(string(out))
	json.Unmarshal(out, result)

	return nil

}
