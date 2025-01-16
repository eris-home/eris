package cli

import (
	"fmt"
	"os"
	"time"
)

type Runner struct {
	InputReader     *InputReader
	OutputFormatter *OutputFormatter
	CommandHandler  *CommandHandler
}

func NewRunner() *Runner {
	runner := &Runner{
		InputReader:     NewInputReader(),
		OutputFormatter: NewOutputFormatter(),
		CommandHandler:  NewCommandHandler(),
	}

	// Register core commands
	runner.CommandHandler.RegisterCommand("help", runner.HelpCommand)
	runner.CommandHandler.RegisterCommand("exit", runner.ExitCommand)
	runner.CommandHandler.RegisterCommand("time", runner.TimeCommand)
	runner.CommandHandler.RegisterCommand("math", runner.MathCommand)

	return runner
}

func (r *Runner) Start() {
	r.OutputFormatter.DisplayInfo("Type 'help' for a list of commands. Type 'exit' to quit.")

	for {
		input := r.InputReader.ReadInput()
		command, args := r.InputReader.ParseInput(input)

		if command == "" {
			r.OutputFormatter.DisplayError("No command entered. Please try again.")
			continue
		}

		r.CommandHandler.ExecuteCommand(command, args)
	}
}

func (r *Runner) HelpCommand(args []string) {
	commands := r.CommandHandler.ListCommands()
	r.OutputFormatter.DisplayInfo(r.OutputFormatter.FormatList(commands))
}

func (r *Runner) ExitCommand(args []string) {
	if r.InputReader.ConfirmAction("Are you sure you want to exit?") {
		r.OutputFormatter.DisplayInfo("Exiting the CLI. Goodbye!")
		os.Exit(0)
	} else {
		r.OutputFormatter.DisplayInfo("Exit cancelled.")
	}
}

func (r *Runner) TimeCommand(args []string) {
	now := time.Now()
	r.OutputFormatter.DisplayInfo(fmt.Sprintf("Current time: %s", now.Format("15:04:05, January 02 2006")))
}

func (r *Runner) MathCommand(args []string) {
	if len(args) < 2 {
		r.OutputFormatter.DisplayError("Please provide two numbers for addition. Example: math 4 5")
		return
	}

	num1 := args[0]
	num2 := args[1]
	r.OutputFormatter.DisplaySuccess(fmt.Sprintf("You entered %s and %s. Let's add them (dummy result: %s+%s)", num1, num2, num1, num2))
}