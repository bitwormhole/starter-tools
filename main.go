package main

import (
	"os"

	"github.com/bitwormhole/starter-tools/tools/configen1"
	"github.com/bitwormhole/starter-tools/tools/configen2"
	"github.com/bitwormhole/starter-tools/tools/configenx"
	"github.com/bitwormhole/starter-tools/tools/help"
)

func main() {
	err := tryMain(os.Args)
	if err == nil {
		return
	}
	panic(err)
}

// return: (command, more...args)
func parseArguments(args []string) (string, []string) {
	command := ""
	more := []string{}
	const CMD_INDEX = 1
	if args != nil {
		for index := range args {
			item := args[index]
			if index == CMD_INDEX {
				command = item
			} else if index > CMD_INDEX {
				more = append(more, item)
			}
		}
	}
	return command, more
}

func tryMain(args []string) error {

	cmd, more := parseArguments(args)
	args = more

	if cmd == "about" {
		return help.RunAbout(args)

	} else if cmd == "configen" {
		return configenx.Run(args)

	} else if cmd == "configen1disable" {
		return configen1.Run(args)

	} else if cmd == "configen2disable" {
		return configen2.Run(args)

	} else if cmd == "help" {
		return help.RunHelpDetail(args)

	} else if cmd == "version" {
		return help.RunVersion(args)

	} else {
		// same as 'help'
		return help.RunHelpDetail(args)
	}
	//	return errors.New("bad command: " + cmd)
}
