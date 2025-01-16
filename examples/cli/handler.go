package cli

import (
	"fmt"
)

type CommandHandler struct {
	Commands map[string]func([]string)
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		Commands: make(map[string]func([]string)),
	}
}

func (ch *CommandHandler) RegisterCommand(name string, handler func([]string)) {
	if _, exists := ch.Commands[name]; exists {
		fmt.Printf("CommandHandler: Command '%s' already exists. Skipping registration.\n", name)
		return
	}
	ch.Commands[name] = handler
	fmt.Printf("CommandHandler: Registered command '%s'\n", name)
}

func (ch *CommandHandler) ExecuteCommand(name string, args []string) {
	handler, exists := ch.Commands[name]
	if !exists {
		fmt.Printf("CommandHandler: Unknown command '%s'. Type 'help' for a list of commands.\n", name)
		return
	}
	fmt.Printf("CommandHandler: Executing command '%s'\n", name)
	handler(args)
}

func (ch *CommandHandler) ListCommands() []string {
	keys := make([]string, 0, len(ch.Commands))
	for key := range ch.Commands {
		keys = append(keys, key)
	}
	return keys
}