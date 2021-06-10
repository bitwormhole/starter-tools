package cmd

type CommandRegistration struct {
	Name    string
	Handler CommandHandler
}
