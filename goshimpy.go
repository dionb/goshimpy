package goshimpy

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

// var SafeMode bool

//Call a python function using the python environment
func Call(name, path string, result interface{}, params ...interface{}) error {

	toPrint, err := buildArgString(params)
	if err != nil {
		log.Println(err.Error())
		return errors.New("Error occurred when trying to build python argument list")
	}

	cmd := exec.Command("python3", "-c", fmt.Sprintf("import json; from %s import %s; args = %s; print(json.dumps(%s(*args)))", path, name, toPrint, name))
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

func buildArgString(args []interface{}) (string, error) {
	argString := ""
	var err error
	var argBytes []byte
	log.Println("received ", len(args), " args")
	switch len(args) {
	case 0:
		argString = "[]"
	case 1:
		argBytes, err = json.Marshal(args[0])
		argString = "[" + string(argBytes) + "]"
	default:
		argBytes, err = json.Marshal(args)
		argString = "json.loads('" + string(argBytes) + "')"
	}
	if err != nil {
		return "", err
	}

	log.Println(argString)

	return argString, nil
}
