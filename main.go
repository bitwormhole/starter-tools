package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/bitwormhole/starter-tools/tools/configen"
)

func main() {
	err := tryMain(os.Args)
	if err == nil {
		return
	}
	panic(err)
}

func tryMain(args []string) error {
	cmd, err := tryGetArgument(args, 1)
	if err != nil {
		return err
	}
	args = args[2:]
	if cmd == "configen" {
		return configen.Run(args)
	}
	return errors.New("bad command: " + cmd)
}

func tryGetArgument(args []string, index int) (string, error) {
	if args == nil {
		return "", errors.New("args==nil")
	}
	size := len(args)
	if index < 0 || index >= size {
		return "", errors.New("index out of args array: " + strconv.Itoa(index))
	}
	return args[index], nil
}
