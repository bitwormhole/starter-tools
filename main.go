package main

import (
	"embed"
	"errors"
	"fmt"
	"os"

	"github.com/bitwormhole/starter-tools/cmd"
	"github.com/bitwormhole/starter-tools/etc"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
)

//go:embed src/main/resources
var resources embed.FS

func main() {

	cb := config.NewBuilderFS(&resources, "src/main/resources")
	cb.SetEnableLoadPropertiesFromArguments(true)
	etc.Config(cb)

	err := tryMain(cb)
	if err != nil {
		panic(err)
	}
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

func exec(context application.Context, cmd_name string, cmd_args []string) error {

	o1, err := context.GetComponent("#commands")
	if err != nil {
		return err
	}

	cmdman, ok := o1.(*cmd.CommandManager)
	if !ok {
		return errors.New("obj.(*cmd.CommandManager): return false")
	}

	return cmdman.ExecuteCommand(cmd_name, cmd_args)
}

func tryMain(cb application.ConfigBuilder) error {

	args := os.Args
	cfg := cb.Create()
	context, err := application.Run(cfg, args)
	if err != nil {
		return err
	}

	cmd, cmdArgs := parseArguments(args)
	err = exec(context, cmd, cmdArgs)
	if err != nil {
		return err
	}

	code, err := application.Exit(context)
	if err != nil {
		return err
	}

	fmt.Println("exit with code ", code)
	return nil
}

/*
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
*/
