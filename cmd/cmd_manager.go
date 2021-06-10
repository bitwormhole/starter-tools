package cmd

import (
	"errors"

	"github.com/bitwormhole/starter/application"
)

type CommandManager struct {
	commands       map[string]*CommandRegistration
	handlerDefault CommandHandler
	context        application.Context
}

func (inst *CommandManager) Inject(context application.Context) error {
	inst.context = context
	return nil
}

func (inst *CommandManager) SetDefaultHandler(h CommandHandler) {
	inst.handlerDefault = h
}

func (inst *CommandManager) Init() error {

	comlist, err := inst.context.GetComponentList("*.scope(singleton)")
	if err != nil {
		return err
	}

	commands := make(map[string]*CommandRegistration)
	for index := range comlist {
		com := comlist[index]
		reg, ok := com.(*CommandRegistration)
		if ok {
			commands[reg.Name] = reg
		}
	}

	inst.commands = commands
	return nil
}

func (inst *CommandManager) GetHandler(cmd string) (CommandHandler, error) {
	handler := inst.handlerDefault
	reg := inst.commands[cmd]
	if reg != nil {
		if reg.Handler != nil {
			handler = reg.Handler
		}
	}
	if handler == nil {
		return nil, errors.New("no handler for command: " + cmd)
	}
	return handler, nil
}

func (inst *CommandManager) ExecuteCommand(cmd string, args []string) error {
	handler, err := inst.GetHandler(cmd)
	if err != nil {
		return err
	}
	return handler.Execute(cmd, args)
}
