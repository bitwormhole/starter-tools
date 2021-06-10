package cmd

import "github.com/bitwormhole/starter/application"

type CommandHandler interface {
	Init(context application.Context) error
	Execute(cmd string, args []string) error
}
